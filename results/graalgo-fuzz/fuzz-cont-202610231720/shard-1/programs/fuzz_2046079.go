package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 int16; F2 int32 }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 uint8 }
func genFn0(p0 uint64) uint {
    p0 = (p0 + p0)
    v5 := GenS1{F0: -1}
    var v6 uint = 310
    _ = p0
    _ = v5
    _ = v6
    return (v6 >> ((2) & 63))
}

func main() {
    v1 := []int{86}
    v2 := make(map[string]string)
    v3 := []string{}
    v4 := []string{"a", "foo bar"}
    switch -191 {
    case -1:
        if true {
            var v7 float32 = ((float32((1.5 / ((1e10 - 1e10) - (0.0 + (float64(3.14159) + (1.5 - 1.5)))))) / float32(1.5)) + 1.5)
            if (360 == 1) {
                for fi8 := 0; fi8 < 15; fi8++ {
                    fmt.Println(v7, fi8)
                    switch (("hello" + "\n\t") + "hello") {
                    case "foo bar":
                        v9 := GenS1{F0: -9223372036854775808}
                        _ = v9
                    }
                    fmt.Println(v7)
                    _ = fi8
                }
            }
            _ = v7
        }
    case 127:
        var v10 uint16 = 42
        _ = v10
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        fmt.Println(fi11)
        _ = fi11
    }
    delete(v2, "a")
    v2["hello"] = "a"
    var v12 float64 = 1.5
    for fi13 := 0; fi13 < 2; fi13++ {
        if true {
            var v14 float64 = 3.14159
            fmt.Println(fi13)
            for fi15 := 0; fi15 < 8; fi15++ {
                if false {
                    v1 = append(v1, 127)
                } else {
                    var v16 int8 = -86
                    _ = v16
                }
                v12 = ((v14 * v14) - v12)
                _ = fi15
            }
            _ = v14
        } else {
            _ = len(v1)
            for fi17 := 0; fi17 < 1; fi17++ {
                var v18 uint16 = 42
                fmt.Println(v12, v18)
                _ = fi17
                _ = v18
            }
        }
        v12 += v12
        v19 := GenS1{F0: 0}
        _ = fi13
        _ = v19
    }
    v20 := GenS1{F0: -1}
    v21 := GenS2{F0: "", F1: 200, F2: -2147483648}
    var v22 int64 = -102
    var v23 uint = 42
    fmt.Println(v22, v23)
    v24 := GenS3{F0: ((("foo bar" + "foo bar") + "") + "\n\t"), F1: 65535}
    for fi25 := 0; fi25 < 15; fi25++ {
        v4 = append(v4, "a")
        v22 -= (^(v22))
        switch ((("hello" + "foo bar") + (("foo bar" + "a") + "foo bar")) + "a") {
        case "a":
            v22 -= -1
        }
        _ = fi25
    }
    v26 := GenS4{F0: uint8(v23)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v26
    _ = v3
    _ = v4
}
