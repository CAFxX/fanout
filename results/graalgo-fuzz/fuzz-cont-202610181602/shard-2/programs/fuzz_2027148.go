package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 uint32 }

func main() {
    var v1 bool = ((("hello" + ("foo bar" + (("\n\t" + ("" + "hello")) + "\n\t"))) + (("" + ((("a" + "") + "a") + "foo bar")) + ("" + "foo bar"))) < ("foo bar" + ""))
    var v2 bool = v1
    var v3 bool = v1
    v4 := make(map[int]int)
    var v5 uint8 = 255
    if (((float32(3.14159) + 1.5) * 0.0) != (0.1 - 0.0)) {
        fmt.Println(v2, v5)
        var v6 string = (("hello" + "") + ("hello" + "a"))
        var v7 float32 = 0.1
        _ = v6
        _ = v7
    }
    var v8 uint32 = 65535
    var v9 uint = 1
    v3 = v3
    v2 = v1
    var v10 uint16 = 1
    v2 = ((false && v2) && true)
    for fi11 := 0; fi11 < 13; fi11++ {
        if v2 {
            fmt.Println(v9)
        }
        for fi12 := 0; fi12 < 7; fi12++ {
            if (!((1 == int16(v8)) || v1)) {
                switch (fi11 % (((fi11 % (((fi11 ^ (fi11 / (((fi12 << ((v9) & 63))) | 1)))) | 1))) | 1)) {
                case 549:
                    v13 := GenS1{F0: -1}
                    _ = v13
                case 4294967295:
                    v5 = uint8(v8)
                }
            }
            _ = fi12
        }
        _ = fi11
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        if (int16(v8) == 1000) {
            v15 := GenS1{F0: 0}
            if v2 {
                var v16 int64 = 42
                var v17 float32 = float32((((((float64(1.5) - 1e10) * (float64(-2.25) + -2.25)) - 0.0) - float64(3.14159)) - 0.0))
                _ = v16
                _ = v17
            }
            fmt.Println(v5, v5)
            _ = v15
        } else {
            v1 = false
            var v18 int16 = 755
            fmt.Println(fi14, v8, v5)
            _ = v18
        }
        for fi19 := 0; fi19 < 3; fi19++ {
            var v20 bool = (v3 && v3)
            _ = fi19
            _ = v20
        }
        if (true && false) {
            v2 = v3
            if v2 {
                fmt.Println(v2, v2)
                var v21 int = fi14
                for fi22 := 0; fi22 < 9; fi22++ {
                    var v23 int64 = -965
                    _ = fi22
                    _ = v23
                }
                _ = v21
            } else {
                switch fi14 {
                case 32767:
                    if (!(651 == 1000)) {
                        for fi24 := 0; fi24 < 4; fi24++ {
                            v25 := v4[-9223372036854775808]
                            delete(v4, (v25 ^ (-(fi24))))
                            var v26 int8 = 127
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                        if v3 {
                            v27 := GenS1{F0: -1}
                            if false {
                                v1 = false
                                for fi28 := 0; fi28 < 14; fi28++ {
                                    v29 := v4[(fi28 | -84)]
                                    v30 := GenS2{F0: -2147483648}
                                    _ = fi28
                                    _ = v29
                                    _ = v30
                                }
                            } else {
                                fmt.Println(v2, v8)
                            }
                            var v31 int8 = 100
                            _ = v27
                            _ = v31
                        }
                        v32 := []int{fi14, fi14, (fi14 >> (((v9 + 2)) & 63))}
                        _ = v32
                    } else {
                        v8 = v8
                        v5 += 255
                    }
                }
                fmt.Println(v5, v2)
                v3 = v1
            }
        } else {
            v33 := GenS3{F0: 42, F1: (v8 + v8)}
            var v34 uint16 = v10
            _ = v33
            _ = v34
        }
        _ = fi14
    }
    v35 := GenS3{F0: -868, F1: v8}
    var v36 int = 613
    v10 += (v10 | v10)
    v8 -= v8
    var v37 string = ""
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
