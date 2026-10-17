package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16 }
func genFn0() int64 {
    var v5 uint8 = 255
    if false {
        switch "\n\t" {
        case "foo bar":
            v5 += v5
        case "":
            v5 = 79
        default:
            fmt.Println(v5)
        }
        switch -9223372036854775808 {
        case 9223372036854775807:
            if (0 <= 2147483647) {
                v5 -= v5
                v5 += v5
                v6 := []int{-75, -1}
                _ = v6
            }
        }
        var v7 float32 = 1.5
        _ = v7
    }
    var v8 uint8 = 42
    v9 := []int{32767, int(v5), 370}
    v5 += v5
    _ = v5
    _ = v8
    _ = v9
    return 42
}

func main() {
    v1 := []int{32767, 32768, 32767}
    v2 := make(map[int]bool)
    v3 := make(map[string]uint32)
    var v4 int8 = 42
    if len(v1) > 0 {
        v1[0] = 32768
    }
    if len(v1) > 0 {
        v1[0] = 68
    }
    v1 = append(v1, 128)
    fmt.Println(v4)
    for fi10 := 0; fi10 < 5; fi10++ {
        v3["\n\t"] = 2147483648
        _ = fi10
    }
    var v11 string = ("\n\t" + ("hello" + ("foo bar" + ((("hello" + "hello") + "hello") + "\n\t"))))
    if (!(true && (2147483647 >= 127))) {
        var v12 float32 = 1e10
        v13 := GenS1{F0: -9223372036854775808}
        if len(v1) > 0 {
            v1[0] = 120
        }
        _ = v12
        _ = v13
    } else {
        switch -638 {
        case 127:
            for fi14 := 0; fi14 < 7; fi14++ {
                v15 := GenS1{F0: fi14}
                v16 := GenS2{F0: 4294967295}
                var v17 uint32 = 1
                _ = fi14
                _ = v15
                _ = v16
                _ = v17
            }
        default:
            if false {
                fmt.Println(v4, v11)
            } else {
                fmt.Println(v4)
            }
        }
    }
    v18 := GenS3{F0: 65535}
    if true {
        v4 += v4
    } else {
        var v19 string = v11
        var v20 int16 = 1
        _ = v19
        _ = v20
    }
    var v21 string = (((("\n\t" + v11) + (v11 + "foo bar")) + (v11 + "")) + "foo bar")
    switch v21 {
    case "foo bar":
        switch -9223372036854775808 {
        case 4294967295:
            var v22 string = (v21 + "")
            _ = v22
        case 9223372036854775807:
            for fi23 := 0; fi23 < 20; fi23++ {
                if true {
                    v4 = v4
                }
                v11 = ("a" + "\n\t")
                var v24 int16 = -561
                _ = fi23
                _ = v24
            }
        }
    }
    if true {
        delete(v2, 32768)
    }
    var v25 uint32 = 2147483647
    v25 = (v25 ^ (v25 & 2))
    v11 = v11
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v18
    _ = v2
    _ = v21
    _ = v25
    _ = v3
    _ = v4
}
