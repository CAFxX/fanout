package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int64 }

func main() {
    var v1 float32 = ((float32(float64(3.14159)) - ((1e10 - (0.1 + 1e10)) / float32(((((1e10 / (3.14159 + float64(1.5))) / (1e10 + (3.14159 - 1e10))) - -2.25) - (3.14159 + 0.0))))) * float32((3.14159 / (-2.25 / (0.1 / 1.0)))))
    v2 := make(map[uint32]string)
    var v3 bool = true
    var v4 string = "a"
    if (v3 || (!v3)) {
        v5 := GenS1{F0: 668, F1: 2, F2: -46}
        var v6 int32 = -2147483648
        _ = v5
        _ = v6
    }
    v3 = v3
    var v7 string = "a"
    switch v4 {
    case "foo bar":
        var v8 bool = true
        _ = v8
    }
    for fi9 := 0; fi9 < 12; fi9++ {
        v3 = (2 != 128)
        _ = fi9
    }
    if v3 {
        v7 = "\n\t"
    }
    if (2147483647 < 0) {
        fmt.Println(v4, v3, v1)
        v3 = v3
        delete(v2, 2147483648)
    }
    if (507 >= -1) {
        switch -1 {
        case 127:
            v10 := GenS1{F0: 1, F1: 1, F2: -1}
            _ = v10
        }
    } else {
        v3 = (("foo bar" + v4) < ((v4 + ((v7 + v4) + "\n\t")) + v7))
        if v3 {
            v2[368] = "hello"
        } else {
            var v11 float64 = 1e10
            for fi12 := 0; fi12 < 15; fi12++ {
                delete(v2, 0)
                var v13 int = (fi12 & fi12)
                _ = fi12
                _ = v13
            }
            _ = v11
        }
    }
    delete(v2, 95)
    for fi14 := 0; fi14 < 1; fi14++ {
        switch "a" {
        case "":
            var v15 uint16 = 255
            _ = v15
        case "hello":
            if ((!v3) || v3) {
                v16 := v2[88]
                _ = v16
            } else {
                v17 := []int{}
                fmt.Println(v7, v3)
                _ = v17
            }
        case "a":
            for fi18 := 0; fi18 < 10; fi18++ {
                v19 := GenS1{F0: 877, F1: 0, F2: 1}
                var v20 uint = 255
                _ = fi18
                _ = v19
                _ = v20
            }
        default:
            fmt.Println(v7, v1)
        }
        if (1 <= 9223372036854775807) {
            var v21 int64 = 0
            v22 := GenS1{F0: 4294967295, F1: 2, F2: v21}
            _ = v21
            _ = v22
        }
        _ = fi14
    }
    v1 = (1e10 / v1)
    var v23 uint32 = 4294967295
    v24 := GenS1{F0: 347, F1: 42, F2: -465}
    v25 := GenS1{F0: 2, F1: 0, F2: -9223372036854775808}
    for fi26 := 0; fi26 < 12; fi26++ {
        v7 = ((v4 + ("hello" + "a")) + ((v4 + ("hello" + "hello")) + "\n\t"))
        v23 += v23
        for fi27 := 0; fi27 < 8; fi27++ {
            v1 = float32(-2.25)
            v7 = ("a" + v7)
            var v28 uint64 = 42
            _ = fi27
            _ = v28
        }
        _ = fi26
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v7
}
