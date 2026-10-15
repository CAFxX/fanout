package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64; F1 int32 }
func genFn0() uint {
    v6 := GenS1{F0: "a"}
    switch ("foo bar" + ("a" + "a")) {
    case "":
        for fi7 := 0; fi7 < 3; fi7++ {
            v8 := []int{(fi7 - fi7), (fi7 - fi7), -9223372036854775808, fi7}
            _ = fi7
            _ = v8
        }
    case "foo bar":
        if (9223372036854775807 >= -2147483648) {
            var v9 int = -9223372036854775808
            for fi10 := 0; fi10 < 17; fi10++ {
                v9 += 32768
                _ = fi10
            }
            v9 = 0
            _ = v9
        }
    case "a":
        switch "hello" {
        case "\n\t":
            switch -333 {
            case -2147483648:
                var v11 bool = (1 == 665)
                _ = v11
            case 800:
                v12 := []int{}
                _ = v12
            default:
                for fi13 := 0; fi13 < 16; fi13++ {
                    if true {
                        v14 := []int{(-(fi13))}
                        if len(v14) >= 2 {
                            v14 = v14[1:len(v14)]
                        }
                        for fi15 := 0; fi15 < 16; fi15++ {
                            v14 = append(v14, fi15)
                            _ = fi15
                        }
                        _ = v14
                    }
                    _ = fi13
                }
            }
        default:
            var v16 uint16 = 65535
            _ = v16
        }
    default:
        switch ("\n\t" + "hello") {
        case "a":
            if ((255 != 42) || false) {
                var v17 uint32 = 42
                v17 = v17
                _ = v17
            }
        case "foo bar":
            var v18 uint32 = 290
            _ = v18
        case "hello":
            var v19 uint8 = 55
            _ = v19
        }
    }
    _ = v6
    return 358
}
func genFn1(p0 uint8, p1 uint64) int32 {
    p1 -= (p1 >> ((286) & 63))
    fmt.Println(p0, p0)
    v20 := []int{868, 32767, -509, -948}
    v21 := GenS1{F0: "\n\t"}
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v20
    _ = v21
    return -80
}

func main() {
    var v1 string = "hello"
    var v2 float64 = 3.14159
    var v3 uint32 = 846
    v4 := make(map[string]string)
    var v5 string = (((v1 + v1) + v1) + ("hello" + (("\n\t" + ("" + "foo bar")) + "\n\t")))
    if (false && (!(true || (v1 > ("a" + v5))))) {
        v22 := GenS1{F0: "hello"}
        switch -9223372036854775808 {
        case 996:
            if true {
                switch 962 {
                case 2147483647:
                    var v23 int16 = 32767
                    _ = v23
                case 0:
                    if true {
                        v24 := GenS2{F0: 390}
                        v1 = v1
                        _ = v24
                    } else {
                        v2 -= v2
                        for fi25 := 0; fi25 < 7; fi25++ {
                            var v26 uint = 2
                            for fi27 := 0; fi27 < 9; fi27++ {
                                v2 -= ((v2 * v2) / v2)
                                var v28 int64 = 5
                                _ = fi27
                                _ = v28
                            }
                            _ = fi25
                            _ = v26
                        }
                        v3 = v3
                    }
                }
                fmt.Println(v3, v3, v2)
            } else {
                fmt.Println(v1)
                switch "hello" {
                case "hello":
                    for fi29 := 0; fi29 < 18; fi29++ {
                        v3 = 42
                        _ = fi29
                    }
                }
            }
        }
        v5 = "hello"
        _ = v22
    } else {
        v30 := GenS3{F0: 18446744073709551615, F1: 556}
        _ = v30
    }
    v2 += (v2 / (v2 - (-2.25 * float64(3.14159))))
    if true {
        v31 := []int{1, 867, 20, 42}
        var v32 int16 = -927
        _ = v31
        _ = v32
    }
    v33 := v4[(v5 + ((v1 + (v1 + v5)) + ""))]
    v34 := GenS3{F0: 1, F1: 1}
    for fi35 := 0; fi35 < 7; fi35++ {
        fmt.Println(v2, v1)
        v36 := v4[v1]
        _ = fi35
        _ = v36
    }
    var v37 int32 = -1
    if (1 <= 700) {
        switch int(v3) {
        case 9223372036854775807:
            for fi38 := 0; fi38 < 5; fi38++ {
                for fi39 := 0; fi39 < 17; fi39++ {
                    var v40 float64 = (v2 + (v2 * (v2 + v2)))
                    v37 = (v37 / ((genFn1(1, 4294967295)) | 1))
                    _ = fi39
                    _ = v40
                }
                v3 = v3
                var v41 float64 = 0.1
                _ = fi38
                _ = v41
            }
        case -663:
            var v42 int8 = 1
            _ = v42
        }
        for fi43 := 0; fi43 < 9; fi43++ {
            var v44 uint32 = 2147483648
            v1 = ((("" + "foo bar") + "hello") + "")
            _ = fi43
            _ = v44
        }
        var v45 int64 = int64(v3)
        _ = v45
    } else {
        if false {
            v3 = v3
        }
        for fi46 := 0; fi46 < 16; fi46++ {
            var v47 float32 = 0.1
            fmt.Println(fi46, v2)
            var v48 int64 = 42
            _ = fi46
            _ = v47
            _ = v48
        }
        var v49 string = (v1 + "")
        _ = v49
    }
    _ = len(v4)
    v50 := []int{4294967295, 32768}
    fmt.Println(v1)
    v33 = v33
    var v51 int16 = -934
    var v52 float32 = (-2.25 - 0.1)
    _ = len(v4)
    switch v33 {
    case "foo bar":
        delete(v4, (v33 + v1))
    case "hello":
        var v53 float64 = v2
        _ = v53
    default:
        fmt.Println(v2)
    }
    var v54 uint8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v50:", len(v50), v50)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v37
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v54
}
