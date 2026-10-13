package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]uint32)
    var v3 string = ((("\n\t" + "foo bar") + "") + "a")
    if (!(-2147483648 > 1)) {
        var v4 int32 = 108
        var v5 int8 = -1
        if (!true) {
            if (!(false && (((float32(0.0) * 0.1) - 1e10) < (1.5 + 3.14159)))) {
                fmt.Println(v3)
            }
            fmt.Println(v5, v4)
        } else {
            if false {
                if false {
                    v4 = int32(v5)
                    v6 := GenS1{F0: 128}
                    _ = v6
                }
            }
            fmt.Println(v5)
        }
        _ = v4
        _ = v5
    }
    v3 = "hello"
    var v7 uint32 = 1
    var v8 uint8 = 152
    fmt.Println(v3, v7, v3)
    if ((0 < 0) && true) {
        if (-503 < 42) {
            v9 := []int{958, -824}
            switch 777 {
            case 32768:
                v10 := GenS1{F0: v8}
                _ = v10
            }
            fmt.Println(v8, v3, v3)
            _ = v9
        }
        v11 := GenS1{F0: (v8 << ((18446744073709551615) & 63))}
        _ = v11
    } else {
        var v12 uint = 706
        var v13 int64 = 1
        for fi14 := 0; fi14 < 18; fi14++ {
            v15 := GenS2{F0: "hello"}
            _ = fi14
            _ = v15
        }
        _ = v12
        _ = v13
    }
    switch 32767 {
    case -181:
        v8 = (v8 & v8)
    case 801:
        var v16 uint16 = 1
        _ = v16
    }
    var v17 float32 = 1e10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
