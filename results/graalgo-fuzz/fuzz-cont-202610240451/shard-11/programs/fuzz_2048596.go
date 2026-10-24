package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int; F1 int }
type GenS4 struct { F0 bool; F1 uint8; F2 int16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 bool; F1 uint; F2 uint }
type GenS8 struct { F0 uint; F1 int32; F2 uint16 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int16; F1 uint16 }
func genFn0() uint16 {
    var v7 int8 = 0
    if ((("a" + "") < ("hello" + "foo bar")) && ((("a" + ("hello" + "\n\t")) + "\n\t") >= "")) {
        v8 := []int{-274, 1, 9223372036854775807}
        _ = v8
    } else {
        switch 1 {
        case 32767:
            v9 := GenS1{F0: 0, F1: -9223372036854775808, F2: 1}
            _ = v9
        }
        v10 := GenS2{F0: "hello"}
        _ = v10
    }
    v11 := []int{-472, 0, -2147483648, 819}
    _ = v11
    _ = v7
    return 1
}

func main() {
    v1 := []int{}
    var v2 int8 = 0
    var v3 uint8 = 71
    var v4 float32 = 1e10
    var v5 int8 = v2
    var v6 bool = (-486 == -32768)
    v2 = 127
    for fi12 := 0; fi12 < 13; fi12++ {
        var v13 float32 = (float32(3.14159) * 0.1)
        _ = fi12
        _ = v13
    }
    v6 = true
    v14 := GenS3{F0: 32767, F1: 1}
    if ((-580 > 4294967295) || (v6 && (255 <= 946))) {
        for fi15 := 0; fi15 < 17; fi15++ {
            for fi16 := 0; fi16 < 19; fi16++ {
                if false {
                    _ = len(v1)
                    var v17 float32 = v4
                    fmt.Println(fi16)
                    _ = v17
                }
                _ = fi16
            }
            v3 = v3
            v1 = append(v1, (fi15 << ((0) & 63)))
            _ = fi15
        }
        v5 = (v5 & (v2 >> ((1) & 63)))
    } else {
        var v18 int8 = -85
        var v19 int16 = 1000
        _ = len(v1)
        _ = v18
        _ = v19
    }
    var v20 uint16 = 1
    if (v6 && v6) {
        var v21 float32 = v4
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v21
    } else {
        var v22 int = 128
        v3 = 2
        _ = v22
    }
    if true {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        for fi23 := 0; fi23 < 4; fi23++ {
            switch (fi23 ^ (fi23 * (fi23 >> ((975) & 63)))) {
            case 0:
                for fi24 := 0; fi24 < 5; fi24++ {
                    if len(v1) > 0 {
                        v1[0] = 9223372036854775807
                    }
                    _ = fi24
                }
            case -2147483648:
                for fi25 := 0; fi25 < 16; fi25++ {
                    v20 = (^(v20))
                    _ = fi25
                }
            case 510:
                var v26 uint16 = (v20 >> ((860) & 63))
                _ = v26
            }
            fmt.Println(v6, v5, v2)
            for fi27 := 0; fi27 < 13; fi27++ {
                if (486 != 2) {
                    if true {
                        var v28 float32 = float32(0.0)
                        v20 -= v20
                        var v29 int8 = v5
                        _ = len(v1)
                        v1 = append(v1, 42)
                        _ = v28
                        _ = v29
                    }
                    if len(v1) > 0 {
                        v1[0] = (fi23 | fi23)
                    }
                    var v30 int64 = 1
                    fmt.Println(v30)
                    var v31 uint32 = uint32(v20)
                    _ = v30
                    _ = v31
                }
                _ = fi27
            }
            _ = fi23
        }
        switch ("hello" + ("" + "")) {
        case "a":
            if (-9223372036854775808 <= int64(v5)) {
                v4 = 0.0
                if v6 {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                _ = len(v1)
            } else {
                v3 += (v3 / ((v3) | 1))
                for fi32 := 0; fi32 < 20; fi32++ {
                    v33 := GenS2{F0: "foo bar"}
                    v6 = (2 >= 1)
                    _ = fi32
                    _ = v33
                }
                v34 := GenS4{F0: v6, F1: 1, F2: -655}
                _ = v34
            }
        }
    }
    v35 := GenS5{F0: 1}
    switch 32767 {
    case 437:
        if v6 {
            var v36 int64 = -319
            v36 = v36
            var v37 int = 4294967295
            _ = v36
            _ = v37
        }
    case 245:
        v5 += v2
    }
    for fi38 := 0; fi38 < 13; fi38++ {
        v39 := GenS3{F0: (fi38 & (-(fi38))), F1: fi38}
        for fi40 := 0; fi40 < 15; fi40++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v41 := GenS6{F0: (v20 / ((65535) | 1))}
            switch "" {
            case "hello":
                var v42 uint16 = v20
                _ = v42
            case "\n\t":
                v5 += 0
            }
            _ = fi40
            _ = v41
        }
        _ = fi38
        _ = v39
    }
    var v43 uint8 = (v3 & (v3 + v3))
    switch 0 {
    case -9223372036854775808:
        v44 := GenS7{F0: (v6 || (127 > -1)), F1: 2, F2: 1}
        _ = v44
    default:
        fmt.Println(v6, v4, v2)
    }
    if (!(true || (v6 || v6))) {
        var v45 uint32 = 849
        v46 := GenS8{F0: 0, F1: -2147483648, F2: (v20 * (^(v20)))}
        var v47 uint8 = v43
        _ = v45
        _ = v46
        _ = v47
    }
    v48 := GenS8{F0: 321, F1: 1, F2: 32767}
    if v6 {
        if ("" >= (("a" + (("a" + "") + "")) + ("" + ""))) {
            v49 := GenS9{F0: true}
            v50 := GenS9{F0: v6}
            _ = v49
            _ = v50
        } else {
            v4 = v4
            for fi51 := 0; fi51 < 3; fi51++ {
                switch 2147483647 {
                case 2147483647:
                    switch 4294967295 {
                    case 128:
                        switch "\n\t" {
                        case "":
                            v20 = v20
                        case "\n\t":
                            var v52 int16 = 760
                            _ = v52
                        }
                    case 0:
                        var v53 bool = (int8(v20) != v5)
                        _ = v53
                    }
                case -2147483648:
                    v20 = v20
                case -1:
                    fmt.Println(fi51, v43, v4)
                }
                if len(v1) > 0 {
                    v1[0] = fi51
                }
                for fi54 := 0; fi54 < 7; fi54++ {
                    v55 := GenS10{F0: 0, F1: (v20 >> ((692) & 63))}
                    _ = fi54
                    _ = v55
                }
                _ = fi51
            }
        }
        var v56 int64 = 137
        _ = v56
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v3
    _ = v35
    _ = v4
    _ = v43
    _ = v48
    _ = v5
    _ = v6
}
