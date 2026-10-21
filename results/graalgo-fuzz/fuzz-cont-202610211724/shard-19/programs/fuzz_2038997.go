package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int16 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]string)
    var v3 float64 = (0.1 * (-2.25 / 1.0))
    v4 := make(map[string]string)
    var v5 uint64 = 18446744073709551615
    fmt.Println(v5)
    switch -553 {
    case 731:
        v5 -= v5
    }
    v5 = 832
    delete(v4, "a")
    v5 -= v5
    _ = len(v2)
    for fi6 := 0; fi6 < 7; fi6++ {
        for fi7 := 0; fi7 < 6; fi7++ {
            var v8 string = "foo bar"
            var v9 int = fi7
            v9 -= fi6
            _ = fi7
            _ = v8
            _ = v9
        }
        switch -9223372036854775808 {
        case 310:
            v5 = v5
        case 42:
            var v10 uint16 = 65535
            _ = v10
        case 72:
            for fi11 := 0; fi11 < 13; fi11++ {
                v3 += v3
                fmt.Println(v3)
                v12 := v4["hello"]
                _ = fi11
                _ = v12
            }
        default:
            v5 = 0
        }
        _ = fi6
    }
    for fi13 := 0; fi13 < 18; fi13++ {
        if false {
            switch ("hello" + "foo bar") {
            case "foo bar":
                var v14 float32 = 0.1
                _ = v14
            case "hello":
                v2["foo bar"] = "a"
            }
            fmt.Println(v3)
        } else {
            _ = len(v4)
        }
        if ((!(("hello" + ("a" + "\n\t")) < "\n\t")) || (((!true) && false) || true)) {
            var v15 string = ("hello" + "")
            _ = v15
        } else {
            fmt.Println(v3)
        }
        v3 = float64(((3.14159 + ((3.14159 + 3.14159) + 0.1)) / 3.14159))
        _ = fi13
    }
    v5 = v5
    var v16 int8 = 100
    if (1 == 0) {
        var v17 uint = 1
        for fi18 := 0; fi18 < 3; fi18++ {
            v3 = v3
            _ = fi18
        }
        var v19 uint32 = 2147483647
        _ = v17
        _ = v19
    } else {
        if (!false) {
            v20 := []int{-2147483648, 1, 4294967295}
            delete(v1, 2147483647)
            v5 -= v5
            _ = v20
        } else {
            v21 := []int{}
            var v22 int32 = -2147483648
            _ = v21
            _ = v22
        }
        for fi23 := 0; fi23 < 8; fi23++ {
            var v24 int64 = 468
            _ = fi23
            _ = v24
        }
    }
    v25 := GenS1{F0: 0, F1: 0, F2: 42}
    if false {
        fmt.Println(v16, v16, v3)
    } else {
        v5 -= (v5 + (v5 + (^(v5))))
    }
    var v26 int64 = 613
    v3 -= (float64(1.5) * v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
