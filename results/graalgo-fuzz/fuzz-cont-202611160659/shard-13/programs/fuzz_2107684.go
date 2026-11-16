package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    v1 := make(map[string]bool)
    var v2 int8 = -64
    v3 := []string{"\n\t", ("hello" + "\n\t"), "hello"}
    var v4 float32 = 1.5
    switch -246 {
    case -214:
        v4 += (v4 - float32((-2.25 + 3.14159)))
    case 32768:
        v3 = append(v3, "")
    case 650:
        v5 := GenS1{F0: 65535}
        _ = v5
    }
    for fi6 := 0; fi6 < 3; fi6++ {
        for fi7 := 0; fi7 < 13; fi7++ {
            v4 -= v4
            var v8 int = fi6
            _ = fi7
            _ = v8
        }
        var v9 int16 = 137
        v3 = append(v3, ("\n\t" + "hello"))
        _ = fi6
        _ = v9
    }
    if (322 >= -195) {
        if (458 > 18446744073709551615) {
            v10 := GenS1{F0: 2147483648}
            _ = v10
        } else {
            var v11 float64 = -2.25
            v11 = (v11 - (((v11 * (0.0 / 3.14159)) + v11) + v11))
            v12 := v1["\n\t"]
            _ = v11
            _ = v12
        }
    } else {
        var v13 string = "foo bar"
        _ = v13
    }
    if (0 < -1) {
        v2 = (^(v2))
        v2 = (v2 - 1)
    }
    v14 := GenS1{F0: 4294967295}
    v15 := GenS1{F0: 976}
    var v16 int64 = 194
    var v17 uint32 = 0
    if false {
        fmt.Println(v17, v17)
    }
    v2 = 42
    delete(v1, "a")
    var v18 int32 = -368
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
