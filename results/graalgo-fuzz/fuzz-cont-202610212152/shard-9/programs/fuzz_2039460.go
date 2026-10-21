package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 int, p1 uint8) int64 {
    if false {
        var v4 uint = uint(p1)
        _ = v4
    }
    v5 := []int{p0, p0}
    _ = p0
    _ = p1
    _ = v5
    return 9223372036854775807
}
func genFn1() uint32 {
    v6 := []int{-2147483648}
    if (590 > 1000) {
        var v7 int = 32767
        _ = v7
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    switch "hello" {
    case "hello":
        var v8 int64 = 1
        _ = v8
    case "foo bar":
        var v9 string = ("a" + (("\n\t" + ("hello" + "\n\t")) + ""))
        _ = v9
    default:
        if len(v6) > 0 {
            v6[0] = 128
        }
    }
    var v10 uint32 = 2
    _ = v10
    _ = v6
    return (v10 | v10)
}

func main() {
    v1 := []int{4294967295, -1, 1}
    var v2 int16 = 1000
    v3 := make(map[string]bool)
    switch "foo bar" {
    case "a":
        _ = len(v3)
    }
    fmt.Println(v2)
    v11 := GenS1{F0: -122}
    v12 := GenS1{F0: v2}
    v2 = 1
    v2 = v2
    v13 := v3[(("" + "foo bar") + "hello")]
    v14 := GenS1{F0: (v2 & (v2 & v2))}
    switch (("" + "a") + ("" + ("\n\t" + "\n\t"))) {
    case "foo bar":
        var v15 int = 32768
        _ = v15
    case "a":
        if (0 > 460) {
            var v16 bool = (29 >= 255)
            _ = v16
        } else {
            for fi17 := 0; fi17 < 10; fi17++ {
                var v18 float64 = ((-2.25 + ((float64(float32(1.5)) / 1e10) * (((-2.25 + 3.14159) * ((3.14159 + 3.14159) - 0.0)) - float64(3.14159)))) + 0.0)
                _ = fi17
                _ = v18
            }
            fmt.Println(v2, v2)
            v19 := GenS1{F0: -1}
            _ = v19
        }
    case "hello":
        v3["\n\t"] = (!false)
    default:
        for fi20 := 0; fi20 < 15; fi20++ {
            v13 = true
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            for fi21 := 0; fi21 < 13; fi21++ {
                v2 += 94
                var v22 uint16 = 32767
                _ = fi21
                _ = v22
            }
            _ = fi20
        }
    }
    for fi23 := 0; fi23 < 7; fi23++ {
        var v24 float32 = float32(3.14159)
        var v25 int = (fi23 / ((fi23) | 1))
        v26 := GenS1{F0: -32768}
        _ = fi23
        _ = v24
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 2; fi27++ {
        fmt.Println(v2)
        switch (-(fi27)) {
        case 622:
            var v28 string = ""
            _ = v28
        default:
            var v29 int = fi27
            _ = v29
        }
        _ = fi27
    }
    v2 = (v2 + (v2 / ((v2) | 1)))
    if len(v1) > 0 {
        v1[0] = 0
    }
    var v30 int64 = genFn0(293, 2)
    for fi31 := 0; fi31 < 8; fi31++ {
        v32 := GenS1{F0: v2}
        var v33 int16 = v2
        var v34 uint32 = 756
        _ = fi31
        _ = v32
        _ = v33
        _ = v34
    }
    var v35 int32 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v30
    _ = v35
}
