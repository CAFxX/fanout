package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 int16 }

func main() {
    var v1 int8 = -118
    v2 := []uint32{2147483648, 42}
    v3 := make(map[string]int)
    var v4 uint8 = 255
    var v5 uint64 = 4294967295
    v6 := make(map[int]string)
    if true {
        switch "a" {
        case "a":
            switch 4294967295 {
            case 128:
                v2 = append(v2, 65535)
            case 8:
                delete(v3, "foo bar")
            }
        case "":
            v1 = (^(v1))
        }
        if false {
            var v7 int32 = 0
            v1 = (v1 >> ((941) & 63))
            var v8 int = -487
            _ = v7
            _ = v8
        }
        switch (((("\n\t" + "hello") + "hello") + ("\n\t" + "foo bar")) + "\n\t") {
        case "foo bar":
            fmt.Println(v5)
        }
    }
    fmt.Println(v1, v1)
    v9 := GenS1{F0: 4294967295, F1: -363, F2: 32767}
    fmt.Println(v4)
    var v10 string = (("a" + (("hello" + "hello") + "\n\t")) + ("" + ""))
    for fi11 := 0; fi11 < 4; fi11++ {
        var v12 uint16 = 916
        _ = fi11
        _ = v12
    }
    var v13 int32 = 1
    if ((!true) || (4294967295 < 1)) {
        fmt.Println(v13, v1, v10)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if (false || true) {
            var v14 uint = 2
            var v15 bool = ((219 > uint32(v14)) || (true && (((443 < 2) && (42 <= 0)) && (false && false))))
            _ = v14
            _ = v15
        }
    }
    v13 -= v13
    var v16 int32 = -1
    fmt.Println(v4, v1, v5)
    v17 := GenS1{F0: 18446744073709551615, F1: 1, F2: 1000}
    var v18 int64 = 42
    v10 = (("" + ((("a" + (("" + "\n\t") + ((v10 + v10) + (v10 + v10)))) + v10) + v10)) + ("hello" + ((v10 + v10) + v10)))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
