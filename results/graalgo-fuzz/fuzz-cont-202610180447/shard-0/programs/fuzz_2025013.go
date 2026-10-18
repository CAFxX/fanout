package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8 }
func genFn0(p0 uint64) uint32 {
    if (870 <= 1) {
        p0 = (p0 >> ((4294967295) & 63))
        v4 := []int{}
        _ = v4
    } else {
        var v5 int32 = -338
        _ = v5
    }
    fmt.Println(p0)
    switch "hello" {
    case "foo bar":
        fmt.Println(p0)
    case "":
        var v6 float32 = 0.0
        _ = v6
    case "a":
        switch 128 {
        case 32768:
            if (true || false) {
                for fi7 := 0; fi7 < 11; fi7++ {
                    fmt.Println(fi7)
                    _ = fi7
                }
            } else {
                p0 += 42
            }
        case 581:
            v8 := []int{128, 9223372036854775807, 2147483647}
            _ = v8
        }
    }
    var v9 string = "\n\t"
    var v10 uint32 = 2147483648
    _ = p0
    _ = v10
    _ = v9
    return 2147483648
}

func main() {
    var v1 uint32 = 4294967295
    var v2 float32 = (1e10 / -2.25)
    v3 := []int{823, 2147483647}
    v11 := GenS1{F0: 127}
    var v12 uint = 0
    switch 4294967295 {
    case 32767:
        if len(v3) > 0 {
            v3[0] = 42
        }
    default:
        for fi13 := 0; fi13 < 9; fi13++ {
            if false {
                var v14 string = (("\n\t" + "\n\t") + ((((("a" + "hello") + "a") + "\n\t") + (((("a" + (((("hello" + "\n\t") + "hello") + (("" + "a") + (("\n\t" + ("foo bar" + "\n\t")) + "foo bar"))) + "\n\t")) + "") + "a") + "a")) + "hello"))
                _ = v14
            } else {
                var v15 uint64 = 1
                v16 := GenS1{F0: fi13}
                _ = v15
                _ = v16
            }
            _ = fi13
        }
    }
    if false {
        if ((v12 / (((v12 ^ v12)) | 1)) >= (v12 % ((v12) | 1))) {
            if (0 >= -329) {
                var v17 float32 = (3.14159 - v2)
                v1 -= genFn0(4294967295)
                var v18 uint16 = 32767
                _ = v17
                _ = v18
            }
            if false {
                if len(v3) > 0 {
                    v3[0] = 2147483647
                }
                v19 := GenS2{F0: -1}
                _ = v19
            } else {
                var v20 uint16 = 32767
                v12 += (v12 * 0)
                var v21 int8 = 42
                _ = v20
                _ = v21
            }
            v3 = append(v3, -386)
        } else {
            for fi22 := 0; fi22 < 19; fi22++ {
                v2 = 0.1
                v2 = v2
                _ = fi22
            }
            for fi23 := 0; fi23 < 9; fi23++ {
                var v24 bool = true
                var v25 bool = (!(fi23 >= fi23))
                _ = fi23
                _ = v24
                _ = v25
            }
            v3 = append(v3, 287)
        }
        v1 = 1
    }
    var v26 float64 = (float64(3.14159) - (0.0 * 1e10))
    var v27 uint16 = 187
    var v28 int = -1
    fmt.Println(v26, v12)
    var v29 int32 = -1
    v3 = append(v3, (v28 & 1))
    v28 = int(v27)
    var v30 float32 = float32(v26)
    for fi31 := 0; fi31 < 11; fi31++ {
        v32 := GenS2{F0: 100}
        v30 -= float32(((v26 * float64((v2 / float32(1e10)))) - v26))
        _ = fi31
        _ = v32
    }
    v26 += (v26 + 1.5)
    if false {
        var v33 int32 = 378
        v27 = 749
        for fi34 := 0; fi34 < 18; fi34++ {
            v3 = append(v3, fi34)
            _ = fi34
        }
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
