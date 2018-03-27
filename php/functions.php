<?php

//php获得指定长度的文本函数(中文算2个字符长度)
function get_text( $str, $len ) {
    $newstr = "";
    $count = 0;        
    for( $i = 0; $i < $len && $count < $len; $i++ ) {            
        if( preg_match("/^[\x{4e00}-\x{9fa5}]$/u", mb_substr($str, $i, 1, 'utf8') ) ){
            $count += 2;
        } else {
            $count++;
        }
        $newstr .= mb_substr($str, $i, 1, 'utf8');
    }        
    return $newstr;
}


//PHP抓取远程图片
function getImage($url,$save_dir='',$filename='',$type=0){
    if(trim($url)==''){
        return array('file_name'=>'','save_path'=>'','error'=>1);
    }
    if(trim($save_dir)==''){
        $save_dir='./';
    }
    if(trim($filename)==''){//保存文件名
        $ext=strrchr($url,'.');
        if($ext!='.gif'&&$ext!='.jpg'){
            return array('file_name'=>'','save_path'=>'','error'=>3);
        }
        $filename=time().$ext;
    }
    if(0!==strrpos($save_dir,'/')){
        $save_dir.='/';
    }
    //创建保存目录
    if(!file_exists($save_dir)&&!mkdir($save_dir,0777,true)){
        return array('file_name'=>'','save_path'=>'','error'=>5);
    }
    //获取远程文件所采用的方法
    if($type){
        $ch=curl_init();
        $timeout=5;
        curl_setopt($ch,CURLOPT_URL,$url);
        curl_setopt($ch,CURLOPT_RETURNTRANSFER,1);
        curl_setopt($ch,CURLOPT_CONNECTTIMEOUT,$timeout);
        $img=curl_exec($ch);
        curl_close($ch);
    }else{
        ob_start();
        readfile($url);
        $img=ob_get_contents();
        ob_end_clean();
    }
    //$size=strlen($img);
    //文件大小
    $fp2=fopen($save_dir.$filename,'a');
    fwrite($fp2,$img);
    fclose($fp2);
    unset($img,$url);
    return array('file_name'=>$filename,'save_path'=>$save_dir.$filename,'error'=>0);
}


function popMessage($msg, $url) {
	$script = "
	<script>
		alert( '$msg' );
		window.location.href = '$url';
	</script>
	";
	echo $script;
	die ();
}

function popMessageBack($msg) {
	$script = "<script>
		alert( '$msg' );
		window.history.back();
	</script>
	";
	echo $script;
	die ();
}


//将object转换成array(递归到底)
function get_object_vars_deep($obj) {
	if (is_object ( $obj )) {
		$obj = get_object_vars ( $obj );
	}
	if (is_array ( $obj )) {
		foreach ( $obj as $key => $value ) {
			$obj [$key] = get_object_vars_deep ( $value );
		}
	}
	return $obj;
}

function isEmpty( $val ) {
	$v = trim( $val );
	if( $v === 0 || $v === '0' ) {
		return false;
	}
	return empty( $v );
}


function getLimit( $pageSize ) {
	$pageIndex = @$_GET[ 'page' ];
	if(empty($pageIndex)) {
		$pageIndex = 1;
	}
	$start = ( $pageIndex - 1 ) * $pageSize;
	return $start . ',' .  $pageSize;
}

function isDateTime($param = '', $format = 'Y-m-d H:i:s') {
	return strtotime( date( $format, strtotime( $param ) ) ) === strtotime( $param );
}

function isInteger( $val ) {
	return preg_match("/^\d+$/", $val );	 
}


function convertToIntArray( $stringArray ) {
	foreach( $stringArray  as $key => $value ) {
		$stringArray[$key] = intval( $value );
	}
	return $stringArray;
}

function jsonFormat($data, $indent = null){
	// json encode
	$data = json_encode($data, JSON_UNESCAPED_UNICODE);
	// 缩进处理
	$ret = '';
	$pos = 0;
	$length = strlen($data);
	$indent = isset($indent)? $indent : '    ';
	$newline = "\n";
	$prevchar = '';
	$outofquotes = true;

	for($i=0; $i<=$length; $i++){

		$char = substr($data, $i, 1);

		if($char=='"' && $prevchar!='\\'){
			$outofquotes = !$outofquotes;
		}elseif(($char=='}' || $char==']') && $outofquotes){
			$ret .= $newline;
			$pos --;
			for($j=0; $j<$pos; $j++){
				$ret .= $indent;
			}
		}

		$ret .= $char;

		if(($char==',' || $char=='{' || $char=='[') && $outofquotes){
			$ret .= $newline;
			if($char=='{' || $char=='['){
				$pos ++;
			}

			for($j=0; $j<$pos; $j++){
				$ret .= $indent;
			}
		}

		$prevchar = $char;
	}

	return $ret;
}

//将数字字符串转换城数字
function convert2RealType( $data ) {
    foreach($data as $key => $val) {
        if(gettype($val) == "array") {
            $data[$key] = convert2RealType($val);
        } else if( is_numeric($val) ) {
            $data[$key] = doubleval($val);
        }
    }
    return $data;
}