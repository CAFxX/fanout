package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0() uint16 {
    var v5 uint = 65535
    for fi6 := 0; fi6 < 7; fi6++ {
        fmt.Println(fi6)
        v7 := GenS1{F0: 2}
        var v8 float64 = 0.1
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 19; fi9++ {
        for fi10 := 0; fi10 < 5; fi10++ {
            switch (-(fi9)) {
            case 0:
                v5 -= (^(v5))
            default:
                fmt.Println(fi9, fi10, fi9)
            }
            fmt.Println(fi9, fi9)
            _ = fi10
        }
        for fi11 := 0; fi11 < 9; fi11++ {
            var v12 uint32 = 129
            var v13 string = ("a" + "hello")
            _ = fi11
            _ = v12
            _ = v13
        }
        for fi14 := 0; fi14 < 3; fi14++ {
            switch "a" {
            case "foo bar":
                var v15 bool = (128 != 194)
                v5 -= v5
                if ("a" > "foo bar") {
                    if (v15 || ((4294967295 == 0) && (!((v15 && (963 >= 65535)) && v15)))) {
                        v5 = (v5 / ((2) | 1))
                        var v16 uint16 = 255
                        var v17 string = ""
                        _ = v16
                        _ = v17
                    }
                    v18 := GenS1{F0: 376}
                    _ = v18
                }
                _ = v15
            case "":
                fmt.Println(fi9)
            case "hello":
                v19 := []int{fi9, 4294967295, (fi14 - fi14), (fi14 + (fi14 + 0))}
                _ = v19
            }
            v20 := GenS1{F0: 42}
            _ = fi14
            _ = v20
        }
        _ = fi9
    }
    fmt.Println(v5)
    _ = v5
    return 65535
}
func genFn1(p0 uint64, p1 uint64) uint64 {
    v21 := GenS1{F0: 42}
    v22 := GenS1{F0: 42}
    var v23 int8 = 86
    _ = p0
    _ = p1
    _ = v21
    _ = v22
    _ = v23
    return (p1 - p1)
}

func main() {
    var v1 int64 = -1
    v2 := []string{}
    v3 := make(map[int]uint32)
    var v4 uint64 = 0
    var v24 float64 = 3.14159
    fmt.Println(v24)
    for fi25 := 0; fi25 < 1; fi25++ {
        _ = len(v3)
        _ = fi25
    }
    switch "foo bar" {
    case "":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v26 int16 = 0
    v27 := GenS1{F0: 255}
    var v28 int16 = (v26 / ((v26) | 1))
    delete(v3, 1)
    v29 := GenS1{F0: 32767}
    if false {
        var v30 uint = 116
        for fi31 := 0; fi31 < 16; fi31++ {
            var v32 int16 = v28
            v32 -= (-(v28))
            for fi33 := 0; fi33 < 16; fi33++ {
                var v34 uint = uint(v4)
                _ = fi33
                _ = v34
            }
            _ = fi31
            _ = v32
        }
        v35 := GenS1{F0: 680}
        _ = v30
        _ = v35
    }
    fmt.Println(v26)
    v4 = (v4 * v4)
    var v36 float64 = (0.1 * ((-2.25 * (((float64(-2.25) - v24) * 1e10) + 3.14159)) * (v24 + -2.25)))
    switch 679 {
    case 32768:
        v36 += v24
    case 42:
        fmt.Println(v36)
    default:
        var v37 int64 = v1
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v36
    _ = v4
}
