package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
func genFn0(p0 string, p1 bool) int64 {
    v7 := []int{32768, 913, 42, 9223372036854775807}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    _ = p0
    _ = p1
    _ = v7
    return 0
}

func main() {
    v1 := make(map[string]bool)
    var v2 bool = false
    v3 := []int{2147483647}
    var v4 float32 = 3.14159
    var v5 uint32 = 1
    var v6 int64 = 0
    var v8 uint64 = 3
    v9 := GenS1{F0: 430, F1: ("a" + ("\n\t" + "\n\t"))}
    var v10 int16 = 694
    var v11 bool = v2
    for fi12 := 0; fi12 < 15; fi12++ {
        var v13 int64 = v6
        var v14 string = "\n\t"
        for fi15 := 0; fi15 < 5; fi15++ {
            var v16 int32 = -2147483648
            var v17 uint64 = (v8 + v8)
            v14 = "foo bar"
            var v18 int32 = (^(v16))
            v19 := v1["hello"]
            _ = fi15
            _ = v16
            _ = v17
            _ = v18
            _ = v19
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    var v20 float64 = -2.25
    var v21 int = int(v10)
    fmt.Println(v21)
    switch (v21 & v21) {
    case 32767:
        switch (v21 >> ((4294967295) & 63)) {
        case 2147483647:
            switch (-(v21)) {
            case -864:
                v5 = v5
            case 2147483647:
                var v22 bool = true
                _ = v22
            default:
                v5 -= v5
            }
        case 174:
            if len(v3) > 0 {
                v3[0] = (v21 << ((291) & 63))
            }
        default:
            if (42 >= 128) {
                _ = len(v3)
            }
        }
    default:
        fmt.Println(v21)
    }
    v23 := GenS1{F0: 42, F1: ""}
    var v24 uint8 = 87
    switch "\n\t" {
    case "a":
        switch (v21 >> ((4294967295) & 63)) {
        case -1:
            var v25 uint64 = v8
            _ = v25
        case 1:
            switch v21 {
            case 32768:
                v26 := GenS1{F0: 255, F1: "a"}
                _ = v26
            }
        default:
            switch v21 {
            case 783:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case 1:
                if len(v3) > 0 {
                    v3[0] = v21
                }
            default:
                switch "foo bar" {
                case "a":
                    fmt.Println(v20, v10, v8)
                case "foo bar":
                    fmt.Println(v24, v10, v4)
                }
            }
        }
    case "hello":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v27 int = (v21 % (((-(v21))) | 1))
    for fi28 := 0; fi28 < 6; fi28++ {
        v29 := GenS1{F0: 0, F1: ""}
        _ = fi28
        _ = v29
    }
    if true {
        switch v27 {
        case 127:
            var v30 int8 = 127
            _ = v30
        default:
            _ = len(v3)
        }
        for fi31 := 0; fi31 < 11; fi31++ {
            v32 := GenS1{F0: 1, F1: "hello"}
            _ = len(v1)
            _ = fi31
            _ = v32
        }
        v33 := GenS1{F0: 255, F1: ("foo bar" + "")}
        _ = v33
    } else {
        var v34 int8 = 127
        _ = v34
    }
    for fi35 := 0; fi35 < 4; fi35++ {
        var v36 uint32 = 0
        if (true && false) {
            v11 = ((v10 / ((32767) | 1)) < v10)
            v20 -= float64(0.1)
            fmt.Println(v4, v24)
        }
        _ = fi35
        _ = v36
    }
    var v37 uint8 = 17
    var v38 int32 = 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v24
    _ = v27
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
