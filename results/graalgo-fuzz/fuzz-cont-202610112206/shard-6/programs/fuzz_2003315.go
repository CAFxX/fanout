package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32 }
func genFn0() int32 {
    var v6 float64 = 3.14159
    var v7 int16 = 33
    for fi8 := 0; fi8 < 16; fi8++ {
        if ((2 > 0) && (0 == 774)) {
            var v9 float64 = (float64(0.0) + 3.14159)
            _ = v9
        } else {
            var v10 int8 = 42
            _ = v10
        }
        _ = fi8
    }
    v11 := []int{}
    _ = v11
    _ = v6
    _ = v7
    return -2147483648
}
func genFn1(p0 bool) uint16 {
    v12 := []int{32767, 32768, 127}
    if p0 {
        for fi13 := 0; fi13 < 14; fi13++ {
            _ = len(v12)
            if len(v12) > 0 {
                v12[0] = (fi13 + fi13)
            }
            _ = fi13
        }
        v14 := GenS1{F0: 135}
        _ = v14
    } else {
        var v15 uint32 = 4294967295
        fmt.Println(p0, p0)
        _ = v15
    }
    _ = p0
    _ = v12
    return 465
}

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = true
    v3 := make(map[int]string)
    var v4 uint32 = 730
    v5 := []string{(("foo bar" + "") + (("hello" + ("\n\t" + "a")) + "")), "hello", "a", ("a" + "a")}
    fmt.Println(v2)
    fmt.Println(v4)
    switch "\n\t" {
    case "a":
        for fi16 := 0; fi16 < 3; fi16++ {
            _ = len(v5)
            _ = fi16
        }
    case "hello":
        v17 := GenS1{F0: 2}
        _ = v17
    case "\n\t":
        var v18 float64 = 3.14159
        _ = v18
    default:
        v3[1] = "hello"
    }
    if v2 {
        for fi19 := 0; fi19 < 8; fi19++ {
            if v2 {
                for fi20 := 0; fi20 < 17; fi20++ {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v4 += (v4 | v4)
                    v2 = ((false || v2) && (!(v2 && false)))
                    _ = fi20
                }
                fmt.Println(v2)
                v4 += v4
            }
            v5 = append(v5, "hello")
            _ = fi19
        }
        v21 := GenS2{F0: v4}
        _ = v21
    }
    v2 = (-2147483648 >= 44)
    for fi22 := 0; fi22 < 18; fi22++ {
        fmt.Println(v2, v2, v4)
        v2 = (65535 == 65535)
        _ = fi22
    }
    v4 = v4
    for fi23 := 0; fi23 < 8; fi23++ {
        switch 0 {
        case -2:
            for fi24 := 0; fi24 < 8; fi24++ {
                var v25 uint = 465
                if (v25 >= 1) {
                    var v26 int32 = 0
                    delete(v3, (fi24 * fi23))
                    var v27 uint8 = 127
                    _ = v26
                    _ = v27
                } else {
                    v28 := GenS1{F0: 42}
                    _ = v28
                }
                switch ("hello" + "foo bar") {
                case "\n\t":
                    v29 := GenS2{F0: v4}
                    _ = v29
                }
                _ = fi24
                _ = v25
            }
        case 621:
            if (0 <= (v4 >> ((42) & 63))) {
                var v30 uint32 = v4
                _ = v30
            } else {
                var v31 int32 = 0
                v31 = genFn0()
                if (30 > 1) {
                    var v32 float64 = (0.0 * 3.14159)
                    for fi33 := 0; fi33 < 18; fi33++ {
                        v34 := GenS2{F0: 2}
                        var v35 uint32 = v4
                        v32 = 1e10
                        _ = fi33
                        _ = v34
                        _ = v35
                    }
                    for fi36 := 0; fi36 < 17; fi36++ {
                        v2 = ("\n\t" == "foo bar")
                        if false {
                            fmt.Println(fi23)
                            v37 := GenS2{F0: 4294967295}
                            fmt.Println(v4, v4, fi36)
                            _ = v37
                        }
                        var v38 uint = 255
                        _ = fi36
                        _ = v38
                    }
                    _ = v32
                } else {
                    for fi39 := 0; fi39 < 1; fi39++ {
                        v4 = (v4 % (((v4 * 4294967295)) | 1))
                        var v40 uint64 = 125
                        v1[v4] = fi39
                        _ = fi39
                        _ = v40
                    }
                    v31 = 684
                }
                _ = v31
            }
        case 9223372036854775807:
            var v41 uint32 = v4
            _ = v41
        default:
            for fi42 := 0; fi42 < 20; fi42++ {
                var v43 float32 = (0.0 / 1.5)
                var v44 int8 = 127
                v45 := GenS2{F0: v4}
                _ = fi42
                _ = v43
                _ = v44
                _ = v45
            }
        }
        v46 := GenS3{F0: 42}
        var v47 int16 = 1000
        _ = fi23
        _ = v46
        _ = v47
    }
    v48 := v3[42]
    var v49 uint32 = v4
    var v50 float32 = (3.14159 / 1.0)
    for fi51 := 0; fi51 < 11; fi51++ {
        var v52 int = -408
        _ = fi51
        _ = v52
    }
    v50 = 1.5
    var v53 string = v48
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v53
}
