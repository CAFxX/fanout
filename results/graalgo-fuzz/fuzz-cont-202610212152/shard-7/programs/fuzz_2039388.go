package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 uint16 }
func genFn0(p0 bool, p1 int64) uint {
    fmt.Println(p1, p1)
    p1 = (p1 ^ (p1 & p1))
    _ = p0
    _ = p1
    return 65535
}
func genFn1(p0 uint64, p1 uint8) bool {
    fmt.Println(p0)
    var v7 string = "\n\t"
    _ = p0
    _ = p1
    _ = v7
    return (1e10 == -2.25)
}

func main() {
    var v1 bool = false
    var v2 int16 = 352
    var v3 int = 0
    var v4 bool = v1
    var v5 string = "foo bar"
    v6 := make(map[string]string)
    if v1 {
        v8 := GenS1{F0: 1, F1: -128}
        v4 = v1
        _ = v8
    } else {
        if ((42 >= -311) && true) {
            var v9 int16 = (v2 % ((v2) | 1))
            v9 -= v2
            _ = v9
        }
    }
    var v10 int64 = 0
    for fi11 := 0; fi11 < 7; fi11++ {
        var v12 uint16 = 255
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 14; fi13++ {
        v10 = v10
        delete(v6, ((v5 + "a") + (("a" + "a") + "a")))
        v14 := GenS1{F0: 0, F1: -128}
        _ = fi13
        _ = v14
    }
    var v15 float64 = float64(3.14159)
    var v16 uint32 = 2147483647
    var v17 uint8 = 253
    v18 := []int{v3, -1, -9223372036854775808, v3}
    for fi19 := 0; fi19 < 19; fi19++ {
        v20 := GenS1{F0: (^(v17)), F1: 127}
        _ = fi19
        _ = v20
    }
    v3 += (v3 % (((v3 / ((v3) | 1))) | 1))
    if false {
        v5 = "a"
        for fi21 := 0; fi21 < 12; fi21++ {
            var v22 uint64 = 1
            _ = len(v18)
            _ = fi21
            _ = v22
        }
        for fi23 := 0; fi23 < 18; fi23++ {
            v10 = v10
            _ = fi23
        }
    } else {
        if false {
            var v24 uint64 = 266
            _ = v24
        } else {
            if v1 {
                if false {
                    v25 := GenS2{F0: -128}
                    _ = v25
                } else {
                    for fi26 := 0; fi26 < 12; fi26++ {
                        switch "foo bar" {
                        case "":
                            v27 := GenS3{F0: 6, F1: 32767}
                            _ = v27
                        case "hello":
                            for fi28 := 0; fi28 < 5; fi28++ {
                                var v29 uint64 = 1
                                switch "\n\t" {
                                case "foo bar":
                                    var v30 int64 = v10
                                    _ = v30
                                case "\n\t":
                                    switch (v3 - 32768) {
                                    case 4294967295:
                                        var v31 bool = (int64(fi28) == -352)
                                        _ = v31
                                    }
                                    _ = len(v18)
                                }
                                v10 = v10
                                fmt.Println(v16, fi26, v10)
                                v15 -= v15
                                _ = fi28
                                _ = v29
                            }
                        case "a":
                            _ = len(v6)
                        }
                        if v4 {
                            v18 = append(v18, (fi26 & (fi26 << ((1) & 63))))
                            v32 := GenS3{F0: 2, F1: 42}
                            _ = v32
                        } else {
                            fmt.Println(v16)
                            if ((v4 || v1) || false) {
                                _ = len(v6)
                                if len(v18) >= 2 {
                                    v18 = v18[1:len(v18)]
                                }
                            }
                        }
                        var v33 bool = (3.14159 <= ((3.14159 - 0.1) / ((1.5 - (1e10 - 0.0)) + 1e10)))
                        _ = fi26
                        _ = v33
                    }
                    var v34 int16 = v2
                    _ = v34
                }
                v15 -= v15
            } else {
                var v35 int64 = v10
                _ = v35
            }
        }
        v3 -= v3
        fmt.Println(v3, v16, v3)
    }
    for fi36 := 0; fi36 < 14; fi36++ {
        v37 := GenS1{F0: 255, F1: 1}
        var v38 int8 = 127
        fmt.Println(v1)
        _ = fi36
        _ = v37
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
