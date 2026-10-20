package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint16) int64 {
    v6 := []int{989, 0, 950}
    p0 = p0
    _ = p0
    _ = v6
    return 9223372036854775807
}
func genFn1(p0 int) int16 {
    p0 = 127
    fmt.Println(p0)
    if (!false) {
        fmt.Println(p0)
        switch 9223372036854775807 {
        case -957:
            p0 -= p0
        case 42:
            var v7 uint8 = 42
            _ = v7
        }
    }
    v8 := GenS1{F0: 1, F1: 0}
    v9 := GenS2{F0: 459}
    _ = p0
    _ = v8
    _ = v9
    return -1
}

func main() {
    var v1 uint8 = 255
    var v2 bool = ((-128 >= 100) || (42 < 1))
    var v3 uint = 408
    var v4 uint32 = 89
    v5 := make(map[string]string)
    v10 := GenS1{F0: (v3 >> ((v3) & 63)), F1: 1}
    for fi11 := 0; fi11 < 6; fi11++ {
        v12 := GenS1{F0: v3, F1: 1}
        fmt.Println(fi11, fi11, v4)
        _ = fi11
        _ = v12
    }
    v4 += 42
    v4 += v4
    v1 = (v1 % (((^(v1))) | 1))
    if (!(127 < -1)) {
        v4 = (v4 >> ((v3) & 63))
    } else {
        v13 := GenS2{F0: 2}
        _ = v13
    }
    v14 := GenS3{F0: 134}
    if v2 {
        for fi15 := 0; fi15 < 18; fi15++ {
            for fi16 := 0; fi16 < 20; fi16++ {
                fmt.Println(v4, v4, v1)
                _ = fi16
            }
            var v17 uint = 42
            _ = fi15
            _ = v17
        }
        v18 := GenS2{F0: 428}
        _ = v18
    } else {
        v3 = (v3 % (((v3 << ((uint(v1)) & 63))) | 1))
        v2 = true
        if false {
            if true {
                for fi19 := 0; fi19 < 12; fi19++ {
                    v2 = v2
                    if v2 {
                        v3 = 4294967295
                    } else {
                        var v20 uint32 = (v4 >> ((4294967295) & 63))
                        if (128 != (fi19 & fi19)) {
                            v21 := v5["\n\t"]
                            for fi22 := 0; fi22 < 17; fi22++ {
                                _ = len(v5)
                                var v23 uint = v3
                                _ = fi22
                                _ = v23
                            }
                            v1 = v1
                            _ = v21
                        } else {
                            var v24 uint16 = uint16(v1)
                            _ = v24
                        }
                        _ = v20
                    }
                    var v25 bool = false
                    _ = fi19
                    _ = v25
                }
            }
            _ = len(v5)
        } else {
            v1 = v1
            if (-128 >= 113) {
                v26 := []int{619}
                v5["hello"] = ""
                v27 := GenS3{F0: -32768}
                _ = v26
                _ = v27
            }
            switch -1 {
            case 386:
                var v28 int8 = -128
                _ = v28
            case 1:
                var v29 int16 = 42
                _ = v29
            }
        }
    }
    v30 := []int{32767, 619, 2147483647}
    fmt.Println(v4, v4, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
