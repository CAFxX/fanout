package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 uint; F1 int8 }
func genFn0() int16 {
    var v6 int = -326
    v7 := []int{}
    _ = v6
    _ = v7
    return 32767
}
func genFn1(p0 bool, p1 int64) int32 {
    v8 := []int{}
    p0 = (true && p0)
    var v9 int64 = p1
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 430
}

func main() {
    v1 := []uint32{42, 2147483648}
    var v2 int64 = 129
    v3 := make(map[string]string)
    var v4 uint8 = 255
    var v5 uint8 = (v4 ^ v4)
    var v10 int32 = -144
    v11 := GenS1{F0: 1000}
    fmt.Println(v10, v5, v4)
    var v12 uint8 = v5
    _ = len(v1)
    for fi13 := 0; fi13 < 19; fi13++ {
        if ((946 == 4294967295) || (false || false)) {
            if (1 > 65535) {
                fmt.Println(v5)
                v10 -= 0
                fmt.Println(v4, v10, v10)
            } else {
                var v14 float64 = 1e10
                var v15 int16 = -32768
                _ = v14
                _ = v15
            }
        }
        _ = fi13
    }
    for fi16 := 0; fi16 < 11; fi16++ {
        v4 -= v4
        v2 = 42
        _ = fi16
    }
    v2 -= (v2 * -634)
    v17 := GenS2{F0: 573, F1: 0}
    for fi18 := 0; fi18 < 1; fi18++ {
        v5 = v12
        if (4294967295 <= 1) {
            v4 -= (v12 % (((v4 + v5)) | 1))
        }
        var v19 uint32 = 65535
        _ = fi18
        _ = v19
    }
    v12 = v12
    for fi20 := 0; fi20 < 2; fi20++ {
        fmt.Println(fi20, fi20)
        for fi21 := 0; fi21 < 6; fi21++ {
            if (549 < 599) {
                v22 := GenS3{F0: 18446744073709551615, F1: 46}
                _ = v22
            } else {
                if (232 != (v12 & 1)) {
                    if (65535 != 2) {
                        fmt.Println(fi21, v10, v4)
                        v5 -= (v4 ^ v5)
                    } else {
                        if (!(588 <= 42)) {
                            if true {
                                var v23 string = ("a" + "")
                                var v24 float64 = 3.14159
                                v10 = (v10 % ((971) | 1))
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = v23
                                _ = v24
                            }
                            fmt.Println(v4)
                        } else {
                            var v25 int = 424
                            var v26 uint64 = 2
                            var v27 string = (((("\n\t" + ("a" + (("\n\t" + "\n\t") + "foo bar"))) + ("hello" + ((("foo bar" + "") + "") + ("hello" + "foo bar")))) + "\n\t") + "")
                            v28 := GenS2{F0: 65535, F1: 0}
                            _ = v25
                            _ = v26
                            _ = v27
                            _ = v28
                        }
                    }
                    var v29 int = fi20
                    _ = v29
                }
                if (4294967295 >= 2) {
                    var v30 uint = 262
                    var v31 int8 = -1
                    _ = v30
                    _ = v31
                }
                v12 += (v5 >> ((4294967295) & 63))
            }
            _ = fi21
        }
        _ = fi20
    }
    for fi32 := 0; fi32 < 9; fi32++ {
        var v33 int16 = -958
        _ = fi32
        _ = v33
    }
    if true {
        var v34 uint32 = 0
        var v35 uint = 255
        v36 := GenS1{F0: int16(v35)}
        _ = v34
        _ = v35
        _ = v36
    } else {
        var v37 uint8 = 255
        v37 = 1
        _ = v37
    }
    v5 = 8
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
