package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint }

func main() {
    v1 := make(map[string]bool)
    v2 := []string{}
    v3 := make(map[string]uint32)
    v4 := GenS1{F0: 957}
    var v5 int = 128
    v6 := GenS2{F0: false}
    switch v5 {
    case 127:
        fmt.Println(v5)
    case 128:
        var v7 uint = 0
        _ = v7
    }
    for fi8 := 0; fi8 < 4; fi8++ {
        v5 += 0
        for fi9 := 0; fi9 < 12; fi9++ {
            v10 := GenS1{F0: fi9}
            _ = len(v2)
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    switch v5 {
    case -723:
        for fi11 := 0; fi11 < 10; fi11++ {
            v12 := GenS1{F0: (fi11 - fi11)}
            delete(v1, "\n\t")
            _ = fi11
            _ = v12
        }
    case 9223372036854775807:
        var v13 int32 = 0
        _ = v13
    case 32768:
        var v14 string = "hello"
        _ = v14
    default:
        v5 = (-(v5))
    }
    if false {
        v5 -= v5
    } else {
        if (680 != -1) {
            if false {
                v15 := GenS2{F0: false}
                _ = v15
            }
        }
    }
    if true {
        v16 := GenS3{F0: 251}
        v5 -= 32767
        _ = v16
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if true {
        for fi17 := 0; fi17 < 5; fi17++ {
            for fi18 := 0; fi18 < 12; fi18++ {
                var v19 bool = (-2147483648 <= -2147483648)
                v5 += (fi17 | (v5 * v5))
                switch fi17 {
                case 32767:
                    v5 -= -9223372036854775808
                case 628:
                    v5 += 32767
                }
                _ = fi18
                _ = v19
            }
            fmt.Println(v5)
            _ = fi17
        }
        var v20 uint16 = 578
        _ = v20
    }
    if false {
        v5 += 128
        v5 = (v5 << ((18446744073709551615) & 63))
        v5 = v5
    }
    switch (("a" + ("\n\t" + "\n\t")) + "\n\t") {
    case "":
        switch ("\n\t" + (("\n\t" + "\n\t") + "a")) {
        case "\n\t":
            v5 = v5
        default:
            var v21 uint64 = 4294967295
            _ = v21
        }
    case "hello":
        v5 += 32768
    case "foo bar":
        var v22 int = v5
        _ = v22
    }
    v5 += (v5 / ((v5) | 1))
    v5 += 0
    v23 := GenS2{F0: ((0.1 - (1.5 - 0.1)) >= 0.1)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
