using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Data.SqlClient;
using System.Data;

namespace DBUtility.SqlServer {
    public class ProPageSQL
    {
        public ProPageSQL()
        {
 
        }

        /// <summary>
        /// 获得记录集(分页)
        /// </summary>
        /// <param name="start">起始索引</param>
        /// <param name="end">结束索引</param>
        /// <param name="TableName">表名</param>
        /// <param name="fldName">字段名</param>
        /// <param name="OrderfldName">排序字段</param>
        /// <param name="OrderType">排序类型</param>
        /// <param name="where"></param>
        /// <returns></returns>
        public DataTable GetdataSet(int start, int end,string TableName,string fldName,string OrderfldName,bool OrderType, string where)
        {
            SqlParameter[] parms = { 
                   new SqlParameter("@startIndex",SqlDbType.Int),
                   new SqlParameter("@endindex",SqlDbType.Int),
                   new SqlParameter("@tblName",SqlDbType.VarChar,255),
                   new SqlParameter("@fldName",SqlDbType.VarChar,255),
                   new SqlParameter("@OrderfldName",SqlDbType.VarChar,255),
                   new SqlParameter("@IsReCount",SqlDbType.Bit),
                   new SqlParameter("@OrderType",SqlDbType.Bit),
                   new SqlParameter("@strWhere",SqlDbType.VarChar,1000)
                               };
            parms[0].Value = start;
            parms[1].Value = end;
            parms[2].Value = TableName;
            parms[3].Value = fldName;
            parms[4].Value = OrderfldName;
            parms[5].Value = 0;
            parms[6].Value = OrderType;
            parms[7].Value = where;
            return DbHelperSQL.RunProcedure("[P_GetPagedReCord]", parms,"tb_NetBar").Tables[0];
        }


        /// <summary>
        /// 获得记录集(分页)
        /// </summary>
        /// <param name="start">起始索引</param>
        /// <param name="end">结束索引</param>
        /// <param name="TableName">表名</param>
        /// <param name="fldName">字段名</param>
        /// <param name="OrderfldName">排序字段</param>
        /// <param name="OrderType">排序类型</param>
        /// <param name="where"></param>
        /// <returns></returns>
        public DataTable GetdataSet(int start, int end, string TableName, string fldName, string OrderfldName, bool OrderType, string where,out int Rcount)
        {
            SqlParameter[] parms = { 
                               new SqlParameter("@startIndex",SqlDbType.Int),
                               new SqlParameter("@endindex",SqlDbType.Int),
                               new SqlParameter("@tblName",SqlDbType.VarChar,255),
                               new SqlParameter("@fldName",SqlDbType.VarChar,255),
                               new SqlParameter("@OrderfldName",SqlDbType.VarChar,255),
                               new SqlParameter("@IsReCount",SqlDbType.Bit),
                               new SqlParameter("@OrderType",SqlDbType.Bit),
                               new SqlParameter("@strWhere",SqlDbType.VarChar,1000)
                               };
            parms[0].Value = start;
            parms[1].Value = end;
            parms[2].Value = TableName;
            parms[3].Value = fldName;
            parms[4].Value = OrderfldName;
            parms[5].Value = 0;
            parms[6].Value = OrderType;
            parms[7].Value = where;
            Rcount = GetAllcount(TableName,fldName,OrderfldName,where);
            return DbHelperSQL.RunProcedure("[P_GetPagedReCord]", parms, "tb_NetBar").Tables[0];
        }

        /// <summary>
        /// 取得记录总数
        /// </summary>
        /// <param name="where"></param>
        /// <returns></returns>
        public int GetAllcount(string TableName,string fldName,string OrderfldName,string where)
        {
            SqlParameter[] parms = { 
                               new SqlParameter("@startIndex",SqlDbType.Int),
                               new SqlParameter("@endindex",SqlDbType.Int),
                               new SqlParameter("@tblName",SqlDbType.VarChar,255),
                               new SqlParameter("@fldName",SqlDbType.VarChar,255),
                               new SqlParameter("@OrderfldName",SqlDbType.VarChar,255),
                               new SqlParameter("@IsReCount",SqlDbType.Bit),
                               new SqlParameter("@OrderType",SqlDbType.Bit),
                               new SqlParameter("@strWhere",SqlDbType.VarChar,1000)
                               };
            parms[0].Value = 0;
            parms[1].Value = 0;
            parms[2].Value = TableName;
            parms[3].Value = fldName;
            parms[4].Value = OrderfldName;
            parms[5].Value = 1;
            parms[6].Value = 0;
            parms[7].Value = where;
            return Convert.ToInt32(DbHelperSQL.RunProcedure("[P_GetPagedReCord]", parms, "tb_NetBar").Tables[0].Rows[0][0]);
        }
    }
}
