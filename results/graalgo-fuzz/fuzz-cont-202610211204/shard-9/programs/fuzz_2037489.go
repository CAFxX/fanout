package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 int8; F1 int64; F2 uint }
func genFn0() uint16 {
    v6 := []int{9223372036854775807}
    for fi7 := 0; fi7 < 20; fi7++ {
        if (false && (255 <= 869)) {
            var v8 int8 = -49
            v9 := GenS1{F0: "a", F1: -32768}
            _ = v8
            _ = v9
        } else {
            if len(v6) > 0 {
                v6[0] = -990
            }
            v10 := GenS2{F0: 127, F1: 65535}
            var v11 uint16 = 32767
            _ = v10
            _ = v11
        }
        v12 := GenS3{F0: 280, F1: 438}
        _ = fi7
        _ = v12
    }
    if false {
        for fi13 := 0; fi13 < 10; fi13++ {
            fmt.Println(fi13)
            fmt.Println(fi13)
            _ = fi13
        }
        v14 := GenS1{F0: "", F1: -646}
        _ = v14
    } else {
        if false {
            if false {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                if (!true) {
                    var v15 uint32 = 0
                    v16 := GenS2{F0: 42, F1: 0}
                    v15 = (v15 + (v15 & (v15 + 586)))
                    _ = v15
                    _ = v16
                } else {
                    var v17 uint = 1
                    v17 -= v17
                    v17 += v17
                    _ = v17
                }
            }
            for fi18 := 0; fi18 < 20; fi18++ {
                var v19 int = (fi18 | 118)
                _ = fi18
                _ = v19
            }
        } else {
            if true {
                switch "" {
                case "hello":
                    var v20 string = (("" + "foo bar") + "")
                    _ = v20
                case "\n\t":
                    var v21 float32 = 0.0
                    _ = v21
                default:
                    if false {
                        var v22 int64 = -9223372036854775808
                        _ = v22
                    } else {
                        var v23 uint32 = 2147483648
                        _ = v23
                    }
                }
                var v24 uint64 = 57
                _ = v24
            } else {
                var v25 int8 = 100
                var v26 uint = 320
                _ = v25
                _ = v26
            }
            var v27 int32 = 0
            fmt.Println(v27)
            _ = v27
        }
    }
    var v28 uint = 4294967295
    _ = v28
    _ = v6
    return 32767
}
func genFn1() int64 {
    var v29 int = 882
    v30 := GenS3{F0: 2147483648, F1: 65535}
    _ = v29
    _ = v30
    return -920
}

func main() {
    var v1 uint16 = 1
    var v2 uint64 = 42
    var v3 int16 = -250
    var v4 int32 = 0
    var v5 uint16 = 0
    switch 801 {
    case -1:
        v4 = 0
    }
    var v31 int8 = 1
    if (!((128 < 127) || false)) {
        for fi32 := 0; fi32 < 13; fi32++ {
            var v33 int16 = v3
            var v34 string = "a"
            _ = fi32
            _ = v33
            _ = v34
        }
        v35 := []int{}
        _ = v35
    } else {
        if false {
            for fi36 := 0; fi36 < 16; fi36++ {
                fmt.Println(v2, v3, v31)
                _ = fi36
            }
            var v37 string = ("a" + ("hello" + ""))
            var v38 int64 = genFn1()
            _ = v37
            _ = v38
        }
        for fi39 := 0; fi39 < 15; fi39++ {
            switch fi39 {
            case 1:
                v31 -= 127
            case -9223372036854775808:
                if true {
                    for fi40 := 0; fi40 < 5; fi40++ {
                        if (!(((v31 | (^(v31))) > (v31 & v31)) && ((164 <= 837) || (fi40 > (fi39 + -9223372036854775808))))) {
                            v41 := GenS3{F0: 906, F1: 0}
                            v1 += v5
                            _ = v41
                        } else {
                            if (false || (!true)) {
                                var v42 int64 = 9223372036854775807
                                v3 = (v3 % ((v3) | 1))
                                _ = v42
                            }
                        }
                        v43 := GenS2{F0: 1, F1: 2}
                        _ = fi40
                        _ = v43
                    }
                    v44 := []int{}
                    _ = v44
                }
            case 0:
                v45 := GenS4{F0: (v31 * (-(v31))), F1: 9223372036854775807, F2: 65535}
                _ = v45
            }
            _ = fi39
        }
        var v46 bool = false
        _ = v46
    }
    var v47 bool = ((!("\n\t" <= "\n\t")) && true)
    if v47 {
        v47 = v47
        v48 := []int{-605}
        if len(v48) > 0 {
            v48[0] = 127
        }
        _ = v48
    }
    v47 = (!(((v3 * v3) <= (v3 << ((1) & 63))) || v47))
    var v49 int32 = (v4 + 0)
    v50 := []int{}
    v49 -= v4
    v1 += (v1 * v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v50:", len(v50), v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v47
    _ = v49
    _ = v5
    _ = v50
}
