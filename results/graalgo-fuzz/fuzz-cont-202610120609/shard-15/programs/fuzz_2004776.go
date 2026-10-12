package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 int16; F1 uint16; F2 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int16; F1 int; F2 string }

func main() {
    v1 := make(map[string]string)
    v2 := []int{32768, 33, -1, 42}
    v3 := make(map[uint32]uint32)
    v4 := []string{}
    var v5 int = 127
    var v6 bool = (1 < 4294967295)
    for fi7 := 0; fi7 < 20; fi7++ {
        if v6 {
            v8 := GenS1{F0: 32767, F1: 314}
            if len(v2) > 0 {
                v2[0] = (fi7 + (v5 << ((4294967295) & 63)))
            }
            switch (("\n\t" + "hello") + "a") {
            case "hello":
                for fi9 := 0; fi9 < 10; fi9++ {
                    v6 = v6
                    v6 = (0 >= 18446744073709551615)
                    v10 := GenS2{F0: 0, F1: 0, F2: 127}
                    _ = fi9
                    _ = v10
                }
            case "\n\t":
                v11 := GenS1{F0: 1, F1: 42}
                _ = v11
            }
            _ = v8
        }
        v5 = (fi7 * v5)
        delete(v1, "\n\t")
        _ = fi7
    }
    v12 := GenS3{F0: 65535}
    for fi13 := 0; fi13 < 4; fi13++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        fmt.Println(v5)
        if v6 {
            v14 := GenS2{F0: 856, F1: 0, F2: 42}
            v15 := GenS4{F0: 169, F1: fi13, F2: "a"}
            var v16 int16 = 1
            _ = v14
            _ = v15
            _ = v16
        } else {
            var v17 int = -1
            _ = v17
        }
        _ = fi13
    }
    v5 = v5
    switch ("a" + ("hello" + "foo bar")) {
    case "foo bar":
        for fi18 := 0; fi18 < 2; fi18++ {
            v5 = (-(v5))
            _ = fi18
        }
    case "hello":
        v5 = v5
    case "":
        var v19 int8 = 127
        _ = v19
    }
    if (-128 > 127) {
        v6 = (32767 <= 1000)
        if (!v6) {
            fmt.Println(v6, v5)
            v6 = (!v6)
        } else {
            var v20 uint64 = 451
            var v21 bool = (594 == -120)
            fmt.Println(v6, v20, v20)
            _ = v20
            _ = v21
        }
        var v22 int16 = -438
        _ = v22
    }
    v23 := v3[2]
    v24 := GenS3{F0: 4294967295}
    var v25 string = ("\n\t" + "\n\t")
    v6 = false
    v5 -= v5
    switch 4294967295 {
    case -873:
        fmt.Println(v25, v6)
    }
    var v26 int8 = 0
    switch v25 {
    case "a":
        var v27 int64 = -973
        _ = v27
    case "hello":
        fmt.Println(v26, v25, v26)
    case "\n\t":
        v26 += 0
    default:
        for fi28 := 0; fi28 < 18; fi28++ {
            v29 := GenS2{F0: 1, F1: 1, F2: 255}
            fmt.Println(v5)
            _ = fi28
            _ = v29
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
