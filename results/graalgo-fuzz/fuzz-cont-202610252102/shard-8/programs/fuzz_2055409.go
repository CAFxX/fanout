package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    v1 := make(map[string]bool)
    var v2 int = -2147483648
    var v3 uint16 = 0
    var v4 int32 = -1
    v3 = v3
    delete(v1, "foo bar")
    if (int(v3) >= 42) {
        v5 := GenS1{F0: -122}
        v6 := v1["foo bar"]
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: 940}
    switch (v2 >> ((847) & 63)) {
    case 42:
        v8 := []int{v2, v2, (v2 / ((int(v4)) | 1)), v2}
        _ = v8
    default:
        switch ("\n\t" + "\n\t") {
        case "":
            v9 := []int{0, (v2 | v2)}
            _ = v9
        }
    }
    if true {
        fmt.Println(v3, v2, v2)
    }
    if true {
        v10 := GenS1{F0: -692}
        if (true && true) {
            v4 -= v4
        } else {
            var v11 uint8 = 42
            _ = v11
        }
        for fi12 := 0; fi12 < 13; fi12++ {
            fmt.Println(fi12)
            var v13 bool = false
            v14 := GenS1{F0: 1}
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = v10
    }
    switch v2 {
    case 932:
        var v15 string = "\n\t"
        _ = v15
    default:
        v16 := GenS1{F0: -123}
        _ = v16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
