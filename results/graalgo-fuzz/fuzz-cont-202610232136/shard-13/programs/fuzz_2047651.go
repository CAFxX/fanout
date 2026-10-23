package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 string = ("" + ("" + "\n\t"))
    var v2 string = v1
    var v3 string = "foo bar"
    var v4 int = -480
    var v5 uint32 = 2147483648
    switch (v3 + (((("foo bar" + ("hello" + v3)) + v2) + (v1 + (v1 + (v3 + v1)))) + v2)) {
    case "a":
        var v6 float32 = 0.0
        _ = v6
    case "":
        switch (v1 + (("a" + v2) + v2)) {
        case "":
            fmt.Println(v3)
        case "hello":
            v3 = v1
        default:
            v7 := []int{(v4 + -2147483648), (v4 / ((614) | 1))}
            _ = v7
        }
    case "\n\t":
        var v8 int = v4
        _ = v8
    }
    v5 -= (v5 | v5)
    var v9 int32 = 1
    for fi10 := 0; fi10 < 13; fi10++ {
        var v11 float32 = (1.5 - 0.1)
        for fi12 := 0; fi12 < 1; fi12++ {
            v13 := []int{2147483647, (fi10 + fi12), v4}
            v5 += 4294967295
            v2 = ((v1 + ("foo bar" + v3)) + ("hello" + "\n\t"))
            switch (v1 + "a") {
            case "\n\t":
                var v14 bool = false
                _ = v14
            default:
                var v15 int32 = (v9 - 1)
                _ = v15
            }
            _ = fi12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    var v16 int16 = 861
    if (!true) {
        for fi17 := 0; fi17 < 11; fi17++ {
            v16 = 1
            for fi18 := 0; fi18 < 4; fi18++ {
                v3 = (((v2 + v3) + ("" + "a")) + "a")
                _ = fi18
            }
            var v19 int16 = (v16 >> ((255) & 63))
            _ = fi17
            _ = v19
        }
        var v20 uint16 = 65535
        v21 := []int{1, -123, v4, v4}
        _ = v20
        _ = v21
    } else {
        if (false && false) {
            var v22 string = (v3 + "hello")
            _ = v22
        }
        var v23 string = v1
        var v24 uint = 65535
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 18; fi25++ {
        v26 := GenS1{F0: 4294967295}
        _ = fi25
        _ = v26
    }
    fmt.Println(v4)
    var v27 int = int(v16)
    for fi28 := 0; fi28 < 8; fi28++ {
        v9 -= (v9 | v9)
        _ = fi28
    }
    v9 = v9
    var v29 uint16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
