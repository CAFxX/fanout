package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 int64, p1 uint16) uint32 {
    var v4 int64 = p0
    fmt.Println(p0)
    v5 := GenS1{F0: true}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 0
}

func main() {
    v1 := make(map[uint32]int)
    var v2 string = ("hello" + (("hello" + ("foo bar" + "a")) + ((("" + "hello") + "hello") + "")))
    var v3 uint64 = 4294967295
    var v6 bool = (((false || (-2.25 != 1e10)) || false) && (4294967295 != 4294967295))
    if v6 {
        v7 := []int{2147483647, 685, 2147483647, 128}
        if ((100 < -114) || (((-1 != 1) && (!v6)) && false)) {
            fmt.Println(v3, v6, v6)
        } else {
            if v6 {
                if v6 {
                    var v8 uint64 = 371
                    var v9 uint32 = 65535
                    _ = v8
                    _ = v9
                }
            } else {
                fmt.Println(v6, v6)
                fmt.Println(v2)
            }
            if (!(!false)) {
                if v6 {
                    var v10 uint8 = 155
                    _ = v10
                }
            } else {
                v3 = v3
            }
            var v11 int8 = 0
            _ = v11
        }
        var v12 int16 = -1
        _ = v12
        _ = v7
    }
    v2 = ("\n\t" + v2)
    var v13 int8 = -37
    v13 = -128
    if (-482 <= 1) {
        v14 := GenS1{F0: (((0 > 0) && v6) || false)}
        _ = v14
    }
    v15 := GenS2{F0: genFn0(9223372036854775807, 32767)}
    var v16 float32 = -2.25
    if v6 {
        if v6 {
            var v17 int = -825
            v6 = (false && (v6 || false))
            _ = v17
        } else {
            switch "hello" {
            case "\n\t":
                v16 += v16
            case "":
                v13 = 127
            }
            v18 := []int{4294967295, -436}
            for fi19 := 0; fi19 < 14; fi19++ {
                var v20 string = "hello"
                _ = fi19
                _ = v20
            }
            _ = v18
        }
        v3 = (v3 & (^(v3)))
        v2 = v2
    } else {
        if ((!v6) || v6) {
            v16 -= ((v16 * (0.1 + v16)) - (float32(0.0) * float32((-2.25 - 1.5))))
        } else {
            v21 := GenS3{F0: 1}
            var v22 uint8 = 182
            _ = v21
            _ = v22
        }
        for fi23 := 0; fi23 < 9; fi23++ {
            v3 = (^(v3))
            for fi24 := 0; fi24 < 10; fi24++ {
                var v25 int32 = 0
                v26 := []int{fi24, fi23}
                _ = fi24
                _ = v25
                _ = v26
            }
            _ = fi23
        }
    }
    var v27 float64 = 1e10
    fmt.Println(v6)
    v6 = ((uint(v3) < 65535) && (930 <= 65535))
    var v28 int16 = -32768
    var v29 uint64 = v3
    delete(v1, 937)
    fmt.Println(v3)
    var v30 float32 = (1e10 * v16)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v6
}
