package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 int; F2 uint16 }
func genFn0(p0 uint8) int64 {
    p0 -= p0
    p0 = (p0 & p0)
    v7 := []int{9223372036854775807}
    if true {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        p0 -= (p0 ^ p0)
        for fi8 := 0; fi8 < 4; fi8++ {
            fmt.Println(p0, fi8)
            var v9 bool = true
            _ = fi8
            _ = v9
        }
    }
    _ = p0
    _ = v7
    return 42
}
func genFn1() uint32 {
    var v10 float64 = (float64(-2.25) + float64(1.5))
    var v11 int32 = 0
    _ = v10
    _ = v11
    return 2147483648
}

func main() {
    var v1 int8 = 127
    var v2 uint64 = 0
    var v3 uint64 = 2
    v4 := make(map[int]string)
    v5 := make(map[string]uint32)
    var v6 int16 = 0
    v2 = (^(v3))
    v12 := GenS1{F0: 367}
    for fi13 := 0; fi13 < 1; fi13++ {
        fmt.Println(v3)
        if ("hello" > ("a" + ("hello" + "hello"))) {
            v14 := GenS2{F0: 65535}
            fmt.Println(v3, v1)
            v15 := []int{}
            _ = v14
            _ = v15
        }
        _ = fi13
    }
    v2 = (v2 & (v2 - 2))
    var v16 uint8 = 47
    var v17 string = ""
    var v18 int16 = (v6 >> ((65535) & 63))
    if false {
        if (!true) {
            fmt.Println(v16)
            switch v17 {
            case "a":
                fmt.Println(v18)
            case "foo bar":
                v19 := GenS1{F0: 42}
                _ = v19
            case "\n\t":
                v2 -= (^(v2))
            }
        } else {
            var v20 string = "a"
            _ = v20
        }
    } else {
        for fi21 := 0; fi21 < 17; fi21++ {
            var v22 int = 9223372036854775807
            for fi23 := 0; fi23 < 19; fi23++ {
                v4[-292] = v17
                _ = fi23
            }
            _ = fi21
            _ = v22
        }
        for fi24 := 0; fi24 < 11; fi24++ {
            switch (fi24 / (((-(fi24))) | 1)) {
            case 256:
                var v25 int64 = int64(v16)
                _ = v25
            case -9223372036854775808:
                v2 += 18446744073709551615
            case 9223372036854775807:
                if (v16 < (v16 - (^(v16)))) {
                    if ((!((!true) || true)) || false) {
                        fmt.Println(v17, v17, v17)
                        v18 -= v6
                    } else {
                        if false {
                            v26 := GenS2{F0: 873}
                            v16 = 128
                            v27 := []int{}
                            _ = v26
                            _ = v27
                        }
                        v28 := GenS1{F0: 65535}
                        for fi29 := 0; fi29 < 12; fi29++ {
                            var v30 uint32 = 4294967295
                            v31 := v5[("" + ((((v17 + v17) + ("foo bar" + v17)) + (v17 + v17)) + v17))]
                            var v32 uint = 997
                            _ = fi29
                            _ = v30
                            _ = v31
                            _ = v32
                        }
                        _ = v28
                    }
                    var v33 bool = ((-(fi24)) == 42)
                    _ = v33
                }
            default:
                var v34 int8 = -97
                _ = v34
            }
            v35 := []int{(fi24 - fi24), (fi24 * 9223372036854775807), (^(fi24)), fi24}
            _ = fi24
            _ = v35
        }
    }
    v36 := GenS3{F0: 789, F1: -13, F2: 2}
    v2 += v2
    v37 := GenS2{F0: 408}
    if false {
        _ = len(v5)
        var v38 string = v17
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
