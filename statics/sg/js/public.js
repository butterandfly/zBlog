// JavaScript Document

$(document).ready(function() {
						   //alert("hi");
						   $('#rightnav>ul>li>a').mouseover(function() {
																	 $(this).parent().addClass('mouseselectbig');
																	 });
						   $('#rightnav>ul>li>a').mouseout(function() {
																	$(this).parent().removeClass('mouseselectbig');
																	});
						   $('#rightnav ul ul a').mouseover(function() {
																	 $(this).parent().addClass('mouseselectsmall');
																	 });
						   $('#rightnav ul ul a').mouseout(function() {
																	$(this).parent().removeClass('mouseselectsmall');
																	});
						   });