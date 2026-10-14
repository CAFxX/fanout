package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8 }
func genFn0(p0 int16, p1 int16) uint64 {
    if (1 > 4294967295) {
        var v7 float64 = 0.1
        var v8 float32 = 1e10
        _ = v7
        _ = v8
    } else {
        v9 := GenS1{F0: 42, F1: -364, F2: 2147483647}
        _ = v9
    }
    v10 := []int{}
    _ = p0
    _ = p1
    _ = v10
    return 4294967295
}
func genFn1(p0 int64) int8 {
    v11 := GenS1{F0: 1, F1: -266, F2: -663}
    p0 = (p0 * 704)
    var v12 int32 = -1
    _ = p0
    _ = v11
    _ = v12
    return 1
}

func main() {
    var v1 int = 42
    var v2 int = 42
    v3 := make(map[string]int)
    var v4 float32 = (0.0 - 1.5)
    v5 := make(map[int]bool)
    var v6 uint64 = 631
    for fi13 := 0; fi13 < 18; fi13++ {
        v1 += fi13
        var v14 string = "foo bar"
        for fi15 := 0; fi15 < 16; fi15++ {
            v4 += v4
            switch v14 {
            case "a":
                v16 := []int{}
                _ = v16
            case "foo bar":
                var v17 float64 = 3.14159
                _ = v17
            }
            v2 -= 2147483647
            _ = fi15
        }
        _ = fi13
        _ = v14
    }
    var v18 bool = (!(((v4 - 0.0) + ((((v4 / (v4 + float32(1.5))) / 3.14159) - v4) + 0.0)) != v4))
    if ((!false) && v18) {
        fmt.Println(v18, v4)
        if (382 > -2147483648) {
            v1 = v2
            v19 := GenS1{F0: -90, F1: -230, F2: 2147483647}
            v20 := GenS2{F0: -128}
            _ = v19
            _ = v20
        } else {
            v4 += float32(0.1)
            if v18 {
                if v18 {
                    if v18 {
                        switch "a" {
                        case "foo bar":
                            var v21 float32 = (v4 - (1.5 * 1.5))
                            _ = v21
                        case "hello":
                            v1 = (v1 % (((v2 | (v1 * v1))) | 1))
                        }
                    }
                } else {
                    var v22 string = "hello"
                    _ = v22
                }
            } else {
                v18 = (!(((0 == int8(v2)) || (1e10 > 1e10)) || v18))
            }
        }
        v18 = true
    } else {
        switch "foo bar" {
        case "":
            v23 := []int{0}
            _ = v23
        case "foo bar":
            for fi24 := 0; fi24 < 5; fi24++ {
                v1 += -1
                v4 = v4
                _ = fi24
            }
        case "hello":
            v6 += (v6 >> ((0) & 63))
        }
        for fi25 := 0; fi25 < 1; fi25++ {
            v2 = 128
            v6 = 18446744073709551615
            _ = fi25
        }
    }
    v2 = v1
    v26 := GenS1{F0: 109, F1: (-(v1)), F2: -1}
    switch 32767 {
    case -9223372036854775808:
        for fi27 := 0; fi27 < 15; fi27++ {
            for fi28 := 0; fi28 < 20; fi28++ {
                var v29 int64 = -1
                v30 := v5[(fi27 ^ (fi27 / (((fi27 + 127)) | 1)))]
                var v31 uint32 = 2147483648
                _ = fi28
                _ = v29
                _ = v30
                _ = v31
            }
            _ = fi27
        }
    case 2147483647:
        var v32 int64 = 1
        _ = v32
    case 1:
        switch v1 {
        case -2147483648:
            var v33 int64 = 42
            _ = v33
        }
    }
    switch (("foo bar" + (("hello" + ("hello" + ("hello" + "hello"))) + "a")) + ("foo bar" + "hello")) {
    case "hello":
        var v34 string = "\n\t"
        _ = v34
    }
    v35 := GenS1{F0: -2147483648, F1: 32767, F2: 855}
    v6 = v6
    var v36 int8 = -1
    switch (^(v2)) {
    case 42:
        v37 := GenS3{F0: genFn1(326)}
        _ = v37
    case 9223372036854775807:
        v38 := GenS3{F0: (v36 | 0)}
        _ = v38
    case 32767:
        v4 = float32(1e10)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
