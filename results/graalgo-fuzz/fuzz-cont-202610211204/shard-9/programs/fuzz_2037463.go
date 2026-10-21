package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    v1 := make(map[string]string)
    var v2 uint = 519
    var v3 uint = 657
    fmt.Println(v2, v3)
    switch "" {
    case "":
        switch "a" {
        case "a":
            v2 += 918
        case "foo bar":
            v4 := []int{127, -1, 965}
            _ = v4
        case "\n\t":
            v5 := []int{32767}
            _ = v5
        }
    }
    switch 495 {
    case 4294967295:
        v3 += v3
    case 39:
        switch "foo bar" {
        case "\n\t":
            var v6 uint64 = 42
            _ = v6
        case "hello":
            v2 -= 4294967295
        case "":
            var v7 int16 = 1000
            _ = v7
        }
    case 127:
        v8 := []int{-1, 128, 221, 509}
        _ = v8
    }
    v2 -= 735
    var v9 uint = 255
    if true {
        var v10 int32 = -1
        _ = v10
    }
    fmt.Println(v2)
    v11 := GenS1{F0: 1000}
    if (!false) {
        v12 := []int{-9223372036854775808, -2147483648}
        fmt.Println(v9)
        _ = v12
    } else {
        switch (("a" + (((("a" + "foo bar") + ((("" + (("\n\t" + (("a" + "foo bar") + ("hello" + "\n\t"))) + "")) + ("a" + "a")) + ("\n\t" + "\n\t"))) + ("" + "hello")) + (("foo bar" + "hello") + (("foo bar" + ("foo bar" + "\n\t")) + ("hello" + "foo bar"))))) + "foo bar") {
        case "hello":
            for fi13 := 0; fi13 < 16; fi13++ {
                v14 := []int{fi13, (fi13 * (-(fi13))), fi13, fi13}
                v9 = v2
                _ = fi13
                _ = v14
            }
        case "foo bar":
            fmt.Println(v3, v2)
        default:
            if false {
                v1["hello"] = "\n\t"
                delete(v1, "hello")
            } else {
                if (!true) {
                    if ((4294967295 != 0) && false) {
                        switch (("\n\t" + ("hello" + "\n\t")) + ("a" + "")) {
                        case "":
                            var v15 int16 = -32768
                            _ = v15
                        }
                        v9 = (v3 << (((v3 * v2)) & 63))
                    } else {
                        var v16 uint8 = 141
                        _ = v16
                    }
                }
            }
        }
    }
    var v17 uint8 = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v9
}
