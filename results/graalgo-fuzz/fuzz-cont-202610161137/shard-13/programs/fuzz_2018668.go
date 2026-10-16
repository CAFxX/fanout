package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 int16; F1 int8; F2 bool }
type GenS3 struct { F0 int; F1 uint }
func genFn0(p0 int16) int {
    p0 += p0
    for fi7 := 0; fi7 < 17; fi7++ {
        fmt.Println(fi7, p0)
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 int64 = -9223372036854775808
            _ = fi8
            _ = v9
        }
        var v10 uint8 = 127
        _ = fi7
        _ = v10
    }
    fmt.Println(p0)
    var v11 float64 = (float64(0.0) * 0.1)
    if (!((("\n\t" + ("a" + (((("foo bar" + (("foo bar" + "") + "a")) + "a") + (("\n\t" + "hello") + "foo bar")) + "a"))) + (("hello" + "") + ("\n\t" + ("hello" + "")))) < (("foo bar" + ("\n\t" + (("" + "hello") + ""))) + "hello"))) {
        if (false && (!(1 <= 2))) {
            var v12 int32 = 42
            _ = v12
        } else {
            var v13 int16 = p0
            for fi14 := 0; fi14 < 20; fi14++ {
                if true {
                    var v15 uint16 = 2
                    _ = v15
                } else {
                    for fi16 := 0; fi16 < 4; fi16++ {
                        v11 = v11
                        v17 := GenS1{F0: 0, F1: 127}
                        _ = fi16
                        _ = v17
                    }
                    v18 := []int{fi14}
                    for fi19 := 0; fi19 < 4; fi19++ {
                        if len(v18) > 0 {
                            v18[0] = 0
                        }
                        v18 = append(v18, fi14)
                        if false {
                            fmt.Println(fi19, p0, fi19)
                            if len(v18) > 0 {
                                v18[0] = 42
                            }
                            if true {
                                p0 -= p0
                            } else {
                                fmt.Println(v11)
                            }
                        }
                        _ = fi19
                    }
                    _ = v18
                }
                v20 := []int{32768, -2147483648, (fi14 >> ((42) & 63))}
                for fi21 := 0; fi21 < 12; fi21++ {
                    var v22 int8 = 124
                    _ = fi21
                    _ = v22
                }
                _ = fi14
                _ = v20
            }
            v23 := GenS2{F0: p0, F1: 1, F2: true}
            _ = v13
            _ = v23
        }
        var v24 bool = false
        var v25 bool = (2 == 42)
        _ = v24
        _ = v25
    } else {
        switch int(p0) {
        case -9223372036854775808:
            var v26 uint8 = 42
            _ = v26
        case -1:
            if true {
                var v27 int8 = 45
                _ = v27
            }
        default:
            switch "hello" {
            case "hello":
                v11 += v11
            }
        }
        for fi28 := 0; fi28 < 13; fi28++ {
            var v29 bool = false
            _ = fi28
            _ = v29
        }
    }
    _ = p0
    _ = v11
    return 2147483647
}

func main() {
    v1 := []uint32{2147483647, 4294967295}
    v2 := []int{-174, 0, -2147483648}
    var v3 int16 = -1
    var v4 int64 = int64(v3)
    var v5 int = 1
    var v6 uint = 664
    v30 := GenS1{F0: 47, F1: 1}
    switch ("" + "") {
    case "a":
        for fi31 := 0; fi31 < 10; fi31++ {
            for fi32 := 0; fi32 < 19; fi32++ {
                v4 = (v4 & (v4 & v4))
                _ = fi32
            }
            v6 = 4294967295
            _ = fi31
        }
    case "":
        switch v5 {
        case -2147483648:
            var v33 uint32 = 2
            _ = v33
        }
    default:
        if (!(false && (841 == 1))) {
            var v34 int64 = v4
            if len(v1) > 0 {
                v1[0] = 42
            }
            v35 := GenS3{F0: (v5 | (v5 | v5)), F1: (v6 | (v6 >> ((4294967295) & 63)))}
            _ = v34
            _ = v35
        } else {
            fmt.Println(v5)
            v4 = 0
            v36 := GenS3{F0: 569, F1: 49}
            _ = v36
        }
    }
    fmt.Println(v3, v5, v5)
    v3 -= int16(v5)
    fmt.Println(v5, v5)
    fmt.Println(v6, v6)
    v4 -= -1
    var v37 int16 = (v3 << ((v6) & 63))
    v38 := GenS2{F0: v37, F1: 100, F2: false}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
