package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint }
func genFn0() int {
    var v7 int32 = -1
    switch ("\n\t" + "") {
    case "hello":
        fmt.Println(v7)
    }
    v8 := []int{-86, 4294967295}
    _ = v7
    _ = v8
    return 127
}

func main() {
    var v1 uint16 = 1
    v2 := make(map[string]uint32)
    var v3 uint32 = 4294967295
    var v4 int = 909
    var v5 uint32 = v3
    var v6 int = v4
    var v9 uint8 = 111
    v10 := GenS1{F0: -118}
    v11 := GenS2{F0: 4294967295}
    v12 := GenS1{F0: -9223372036854775808}
    var v13 uint = 2
    if ((!(!false)) || (-128 < 42)) {
        v2["hello"] = v5
        if (!true) {
            var v14 int64 = 691
            var v15 int32 = 2147483647
            var v16 uint16 = uint16(v9)
            _ = v14
            _ = v15
            _ = v16
        } else {
            if (((0 == 42) || (2 < v13)) && false) {
                v6 = v4
                v17 := []int{-431, (v6 << ((uint(v1)) & 63))}
                _ = v17
            } else {
                fmt.Println(v5)
                v18 := GenS2{F0: 4294967295}
                _ = v18
            }
            var v19 int64 = -1
            _ = v19
        }
        var v20 string = ""
        _ = v20
    }
    var v21 int64 = 0
    v13 += 65535
    var v22 string = "foo bar"
    var v23 bool = false
    v6 = -1
    v13 = (v13 % ((v13) | 1))
    var v24 uint = v13
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
