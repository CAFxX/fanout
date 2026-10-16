package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 uint64; F1 int16 }

func main() {
    var v1 uint8 = 128
    v2 := []int{}
    v3 := []string{}
    var v4 bool = false
    switch "hello" {
    case "":
        var v5 int32 = -663
        _ = v5
    case "foo bar":
        v6 := GenS1{F0: 0, F1: v4}
        _ = v6
    default:
        switch "hello" {
        case "foo bar":
            var v7 int32 = 5
            _ = v7
        default:
            var v8 bool = (42 >= 1)
            _ = v8
        }
    }
    var v9 float64 = 1.5
    fmt.Println(v1)
    for fi10 := 0; fi10 < 16; fi10++ {
        var v11 uint8 = v1
        for fi12 := 0; fi12 < 4; fi12++ {
            var v13 int16 = 1000
            _ = len(v2)
            if ((v4 || true) || (100 < int8(fi12))) {
                if len(v2) > 0 {
                    v2[0] = -9223372036854775808
                }
                switch (fi10 / ((int(v13)) | 1)) {
                case 128:
                    _ = len(v3)
                case 42:
                    v11 -= 166
                }
                var v14 uint32 = 519
                _ = v14
            }
            _ = fi12
            _ = v13
        }
        var v15 uint = 0
        _ = fi10
        _ = v11
        _ = v15
    }
    var v16 int64 = -181
    var v17 float32 = 0.0
    if false {
        for fi18 := 0; fi18 < 16; fi18++ {
            v1 += 106
            var v19 float64 = (v9 / 1.0)
            _ = fi18
            _ = v19
        }
        if (v17 == -2.25) {
            switch -137 {
            case 178:
                var v20 int32 = 2147483647
                _ = v20
            case 32768:
                var v21 int = -456
                _ = v21
            default:
                v22 := GenS2{F0: 1, F1: 32767}
                _ = v22
            }
            var v23 int = 1
            v17 += (0.1 * v17)
            _ = v23
        }
        if ("hello" != "a") {
            fmt.Println(v4)
            for fi24 := 0; fi24 < 16; fi24++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                if len(v3) > 0 {
                    v3[0] = ""
                }
                _ = fi24
            }
            v4 = v4
        }
    }
    _ = len(v2)
    var v25 uint64 = 18446744073709551615
    if (0.1 == 0.1) {
        switch 2147483647 {
        case -2147483648:
            v1 = (v1 << ((255) & 63))
        case 2147483647:
            v25 = 46
        }
    } else {
        fmt.Println(v25, v1)
        if (true || (v17 >= (v17 * (v17 * v17)))) {
            var v26 int16 = 0
            v25 = v25
            switch "hello" {
            case "hello":
                v27 := GenS1{F0: (^(v25)), F1: (!v4)}
                _ = v27
            }
            _ = v26
        }
    }
    if len(v2) > 0 {
        v2[0] = -561
    }
    var v28 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v25
    _ = v28
    _ = v3
    _ = v4
    _ = v9
}
