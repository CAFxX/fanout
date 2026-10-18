package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string }
func genFn0(p0 bool) uint32 {
    v4 := GenS1{F0: (4294967295 != 380)}
    for fi5 := 0; fi5 < 15; fi5++ {
        v6 := []int{fi5, fi5}
        _ = fi5
        _ = v6
    }
    p0 = p0
    var v7 int = 32767
    v8 := GenS1{F0: (p0 && true)}
    _ = p0
    _ = v4
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 uint8 = 1
    v2 := make(map[string]uint32)
    v3 := make(map[string]string)
    for fi9 := 0; fi9 < 12; fi9++ {
        var v10 int32 = 42
        switch (("\n\t" + ("hello" + ("" + "hello"))) + "foo bar") {
        case "hello":
            v1 -= 0
        case "\n\t":
            if ("hello" == ((((("foo bar" + "a") + "hello") + ("" + (("a" + ("hello" + "a")) + "foo bar"))) + "foo bar") + "foo bar")) {
                for fi11 := 0; fi11 < 20; fi11++ {
                    var v12 int8 = 127
                    _ = fi11
                    _ = v12
                }
                v13 := GenS1{F0: true}
                _ = v13
            } else {
                var v14 uint = 18446744073709551615
                for fi15 := 0; fi15 < 11; fi15++ {
                    v10 = 890
                    _ = fi15
                }
                v16 := []int{(^(fi9))}
                _ = v14
                _ = v16
            }
        }
        _ = fi9
        _ = v10
    }
    var v17 int64 = 0
    v18 := []int{-768, -2147483648, -637, 4294967295}
    v1 -= v1
    var v19 int32 = 2147483647
    for fi20 := 0; fi20 < 2; fi20++ {
        var v21 uint16 = 685
        _ = fi20
        _ = v21
    }
    if (0 != (v17 | -1)) {
        v3["a"] = "foo bar"
        v22 := GenS2{F0: ""}
        if (4294967295 >= 65535) {
            var v23 int64 = v17
            var v24 float32 = (1.5 / 1.0)
            var v25 int8 = 74
            _ = v23
            _ = v24
            _ = v25
        }
        _ = v22
    } else {
        var v26 int = 127
        var v27 uint = 2
        v19 = (-(v19))
        _ = v26
        _ = v27
    }
    var v28 bool = (-234 == 42)
    if v28 {
        v28 = v28
        var v29 uint8 = (v1 >> ((4294967295) & 63))
        var v30 uint32 = 862
        _ = v29
        _ = v30
    }
    var v31 int32 = (v19 << ((432) & 63))
    v31 -= 484
    switch "\n\t" {
    case "hello":
        if v28 {
            v17 += v17
        }
    case "\n\t":
        var v32 float64 = 1.5
        _ = v32
    case "foo bar":
        for fi33 := 0; fi33 < 5; fi33++ {
            for fi34 := 0; fi34 < 12; fi34++ {
                var v35 int64 = (v17 - v17)
                fmt.Println(v1, fi34, v35)
                if ("a" < "") {
                    v35 -= 670
                } else {
                    v35 = -9223372036854775808
                }
                switch (fi33 - (^(fi34))) {
                case 9223372036854775807:
                    v28 = (((0.1 * (float64(3.14159) - (((1.5 + ((0.1 * (-2.25 / float64(-2.25))) - float64((((((((1e10 * 1e10) + float32((1e10 + 0.1))) - 0.1) + 0.0) - (0.1 / 1.0)) / 1.0) - 0.0)))) / 3.14159) - float64((1.5 - 1e10))))) + (1.5 + (((float64(1e10) * -2.25) * float64(0.0)) * 0.1))) == -2.25)
                case -478:
                    v36 := GenS1{F0: v28}
                    v19 += -2147483648
                    fmt.Println(fi34, v31)
                    _ = v36
                }
                _ = fi34
                _ = v35
            }
            _ = fi33
        }
    }
    v31 = -2147483648
    if len(v18) > 0 {
        v18[0] = -2147483648
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v28
    _ = v3
    _ = v31
}
