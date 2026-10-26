package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    v1 := []int{}
    var v2 float64 = ((0.0 * (3.14159 + 0.1)) / float64(1e10))
    var v3 int32 = 2147483647
    v4 := make(map[int]string)
    switch "hello" {
    case "":
        v3 = v3
    case "\n\t":
        switch 42 {
        case 1:
            if false {
                if false {
                    fmt.Println(v2)
                    fmt.Println(v2)
                }
                v5 := GenS1{F0: "\n\t"}
                fmt.Println(v3)
                _ = v5
            }
        }
    default:
        for fi6 := 0; fi6 < 1; fi6++ {
            v4[4294967295] = "hello"
            if len(v1) > 0 {
                v1[0] = 1
            }
            var v7 int64 = 0
            _ = fi6
            _ = v7
        }
    }
    _ = len(v1)
    v8 := GenS1{F0: (((("" + ("" + "hello")) + "") + "\n\t") + "foo bar")}
    v3 = (v3 >> ((42) & 63))
    var v9 int32 = (v3 % ((v3) | 1))
    var v10 int = 0
    switch (v10 / (((v10 * v10)) | 1)) {
    case 11:
        switch ("hello" + "") {
        case "foo bar":
            v11 := GenS1{F0: "a"}
            _ = v11
        case "a":
            v3 += -2147483648
        }
    case 0:
        fmt.Println(v3, v10, v3)
    default:
        var v12 uint = 2
        _ = v12
    }
    v13 := GenS1{F0: "\n\t"}
    var v14 uint16 = 0
    fmt.Println(v9, v10)
    if (("foo bar" + "\n\t") > "") {
        var v15 int32 = (v9 << ((826) & 63))
        _ = v15
    } else {
        if (v3 != (v9 - v3)) {
            fmt.Println(v10, v9)
            var v16 int32 = v3
            _ = v16
        }
        v17 := GenS1{F0: ((("" + "") + "\n\t") + "foo bar")}
        _ = v17
    }
    var v18 uint64 = 148
    delete(v4, v10)
    var v19 int64 = 0
    if (((^(v14)) >= (v14 * 2)) && (true && ((v2 <= ((float64(1.5) * (v2 / v2)) - ((v2 / 1.0) * 0.1))) || ("\n\t" > "")))) {
        v1 = append(v1, (v10 | -2147483648))
    }
    switch (v10 & v10) {
    case 42:
        v20 := GenS1{F0: "hello"}
        _ = v20
    case 1:
        v9 = v9
    case 128:
        if (1 != 42) {
            for fi21 := 0; fi21 < 13; fi21++ {
                v10 = v10
                _ = len(v1)
                var v22 uint16 = v14
                _ = fi21
                _ = v22
            }
            switch "\n\t" {
            case "\n\t":
                var v23 float64 = (v2 * (v2 * 0.0))
                _ = v23
            case "foo bar":
                v10 = 9223372036854775807
            case "a":
                if (!true) {
                    v4[9223372036854775807] = ""
                    v3 -= -1
                }
            }
            v9 = v3
        }
    }
    fmt.Println(v9)
    var v24 uint32 = 764
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
