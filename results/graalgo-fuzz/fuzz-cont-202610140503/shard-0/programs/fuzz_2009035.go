package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 uint) uint16 {
    p0 += 583
    var v5 int8 = -14
    fmt.Println(p0, p0)
    v5 -= (v5 & v5)
    v6 := GenS1{F0: 0}
    _ = p0
    _ = v5
    _ = v6
    return 255
}

func main() {
    var v1 float32 = 0.1
    var v2 float64 = 3.14159
    var v3 int = 32767
    v4 := []int{}
    v7 := GenS1{F0: genFn0(371)}
    v8 := GenS1{F0: 255}
    for fi9 := 0; fi9 < 5; fi9++ {
        if false {
            v10 := GenS1{F0: 42}
            for fi11 := 0; fi11 < 2; fi11++ {
                fmt.Println(fi9, fi9)
                _ = fi11
            }
            fmt.Println(v2, fi9)
            _ = v10
        }
        if true {
            fmt.Println(v3, v2, v1)
            fmt.Println(v1)
        }
        fmt.Println(fi9, v1)
        _ = fi9
    }
    var v12 bool = true
    v13 := GenS1{F0: 327}
    if (100 < -128) {
        if v12 {
            _ = len(v4)
            for fi14 := 0; fi14 < 20; fi14++ {
                var v15 float32 = ((1.5 - float32(v2)) * 1.5)
                _ = fi14
                _ = v15
            }
        } else {
            var v16 uint = 42
            _ = v16
        }
        v1 -= 0.1
        var v17 uint64 = 2
        _ = v17
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v18 float32 = 1e10
    if (1 <= 4294967295) {
        if true {
            v18 += (v18 - (v1 - v18))
            for fi19 := 0; fi19 < 9; fi19++ {
                fmt.Println(v2, fi19, v18)
                var v20 float32 = 0.0
                _ = fi19
                _ = v20
            }
        }
        fmt.Println(v2)
        if (0.1 >= -2.25) {
            if true {
                v1 += 0.1
                v3 -= v3
            }
            for fi21 := 0; fi21 < 17; fi21++ {
                fmt.Println(v1, v12)
                var v22 int64 = -30
                _ = fi21
                _ = v22
            }
            if (("foo bar" + "foo bar") < "") {
                var v23 int16 = 42
                if len(v4) > 0 {
                    v4[0] = (v3 >> ((4294967295) & 63))
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = v23
            }
        } else {
            var v24 int32 = -2147483648
            var v25 uint64 = 1
            var v26 int16 = 186
            _ = v24
            _ = v25
            _ = v26
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
