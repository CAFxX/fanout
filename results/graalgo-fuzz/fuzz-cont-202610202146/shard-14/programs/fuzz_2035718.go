package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0() string {
    v7 := []int{}
    if false {
        var v8 int8 = -1
        _ = v8
    } else {
        var v9 int8 = 0
        for fi10 := 0; fi10 < 19; fi10++ {
            var v11 float64 = -2.25
            _ = fi10
            _ = v11
        }
        for fi12 := 0; fi12 < 2; fi12++ {
            var v13 float32 = 1e10
            var v14 uint64 = 2
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = v9
    }
    v7 = append(v7, -2147483648)
    var v15 uint16 = 65535
    _ = v15
    _ = v7
    return "hello"
}
func genFn1() int16 {
    var v16 uint16 = 426
    var v17 float64 = 1.5
    switch ("\n\t" + "") {
    case "\n\t":
        v18 := []int{0, 296}
        _ = v18
    }
    v19 := GenS1{F0: 127}
    _ = v16
    _ = v17
    _ = v19
    return -65
}

func main() {
    v1 := []int{32768, 909, 127}
    var v2 float32 = (float32(1e10) - (((3.14159 * 1.5) * 3.14159) / 1.0))
    var v3 string = (("\n\t" + (("a" + "a") + "\n\t")) + (("" + "a") + ""))
    var v4 bool = (((3.14159 * 1.5) * (0.0 + 0.0)) > float64(1.5))
    var v5 int8 = 1
    v6 := make(map[uint32]int)
    v3 = ("hello" + "foo bar")
    fmt.Println(v2, v2, v3)
    fmt.Println(v4, v4)
    var v20 int32 = 42
    v3 = (("foo bar" + "a") + (v3 + ""))
    if (v4 || v4) {
        v3 = v3
    }
    switch v3 {
    case "":
        if len(v1) > 0 {
            v1[0] = int(v20)
        }
    case "foo bar":
        v5 += -128
    case "hello":
        fmt.Println(v5, v4, v3)
    default:
        var v21 bool = v4
        _ = v21
    }
    if (1 != 128) {
        v2 = 1.5
    }
    _ = len(v1)
    v2 += float32(3.14159)
    var v22 int16 = -1
    for fi23 := 0; fi23 < 19; fi23++ {
        for fi24 := 0; fi24 < 5; fi24++ {
            if (v4 && v4) {
                v25 := v6[2]
                _ = v25
            } else {
                var v26 int8 = v5
                delete(v6, 65535)
                _ = len(v1)
                _ = v26
            }
            var v27 uint = 42
            if v4 {
                if (!v4) {
                    var v28 float64 = 1.5
                    switch fi24 {
                    case 4294967295:
                        _ = len(v6)
                    case 9223372036854775807:
                        v6[1] = fi24
                    }
                    _ = v28
                } else {
                    if v4 {
                        var v29 uint16 = 255
                        _ = v29
                    }
                    v6[42] = fi24
                    if false {
                        fmt.Println(fi23, v27, v2)
                    } else {
                        var v30 uint16 = 42
                        _ = v30
                    }
                }
                if (0 < 934) {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v27 = (^(v27))
                    v5 -= v5
                }
                v4 = false
            } else {
                if false {
                    var v31 float64 = 3.14159
                    fmt.Println(v22, v27, v5)
                    _ = v31
                } else {
                    if len(v1) > 0 {
                        v1[0] = (fi24 >> (((^(v27))) & 63))
                    }
                    var v32 uint32 = 65535
                    _ = v32
                }
            }
            v4 = v4
            _ = fi24
            _ = v27
        }
        _ = fi23
    }
    for fi33 := 0; fi33 < 2; fi33++ {
        v34 := GenS1{F0: v5}
        v2 = v2
        for fi35 := 0; fi35 < 8; fi35++ {
            var v36 bool = (2 >= 289)
            var v37 int = (fi33 / ((32767) | 1))
            v22 -= (v22 % (((^(v22))) | 1))
            _ = fi35
            _ = v36
            _ = v37
        }
        _ = fi33
        _ = v34
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
