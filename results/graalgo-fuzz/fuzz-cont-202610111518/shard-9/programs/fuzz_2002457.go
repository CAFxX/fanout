package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int }
func genFn0(p0 uint, p1 int32) uint64 {
    p0 = p0
    var v4 uint64 = 2
    _ = p0
    _ = p1
    _ = v4
    return 971
}

func main() {
    var v1 string = "\n\t"
    var v2 int = 242
    v3 := []uint32{}
    for fi5 := 0; fi5 < 13; fi5++ {
        switch 128 {
        case 32767:
            switch fi5 {
            case 2147483647:
                var v6 float32 = 3.14159
                _ = v6
            default:
                var v7 bool = false
                _ = v7
            }
        case 0:
            var v8 float64 = 3.14159
            _ = v8
        default:
            var v9 uint64 = 42
            _ = v9
        }
        fmt.Println(fi5)
        _ = fi5
    }
    for fi10 := 0; fi10 < 14; fi10++ {
        switch -2147483648 {
        case 32768:
            if true {
                fmt.Println(v2)
            }
        case 744:
            var v11 uint32 = 2147483648
            _ = v11
        case 0:
            _ = len(v3)
        }
        fmt.Println(v1, v1, v2)
        _ = fi10
    }
    switch ("foo bar" + "a") {
    case "hello":
        v2 -= v2
    case "":
        for fi12 := 0; fi12 < 8; fi12++ {
            fmt.Println(v1, fi12)
            switch ("hello" + "\n\t") {
            case "\n\t":
                v3 = append(v3, 65535)
            case "foo bar":
                var v13 int = v2
                _ = v13
            default:
                v2 = -2147483648
            }
            _ = fi12
        }
    case "foo bar":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    if true {
        if (!true) {
            var v14 uint8 = 5
            _ = v14
        }
    } else {
        var v15 bool = true
        var v16 int8 = -6
        _ = v15
        _ = v16
    }
    if len(v3) > 0 {
        v3[0] = 2
    }
    var v17 uint64 = 1
    v1 = ("\n\t" + v1)
    switch (v1 + "\n\t") {
    case "":
        v18 := GenS1{F0: 42}
        _ = v18
    case "hello":
        v17 = genFn0(uint(v17), 0)
    }
    var v19 uint32 = 2
    var v20 string = v1
    if (!false) {
        var v21 uint16 = 406
        var v22 uint64 = 18446744073709551615
        var v23 string = "a"
        _ = v21
        _ = v22
        _ = v23
    } else {
        for fi24 := 0; fi24 < 17; fi24++ {
            for fi25 := 0; fi25 < 3; fi25++ {
                v2 = (fi24 + v2)
                var v26 uint = 65535
                v1 = "\n\t"
                _ = fi25
                _ = v26
            }
            v27 := GenS2{F0: (fi24 << ((65535) & 63))}
            for fi28 := 0; fi28 < 2; fi28++ {
                v20 = v20
                var v29 bool = true
                var v30 float32 = -2.25
                var v31 uint16 = 1
                _ = fi28
                _ = v29
                _ = v30
                _ = v31
            }
            _ = fi24
            _ = v27
        }
    }
    fmt.Println(v19)
    var v32 int64 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v32
}
