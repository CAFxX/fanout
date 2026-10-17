package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }

func main() {
    var v1 int16 = 109
    v2 := []int{}
    v3 := []string{("hello" + ""), "hello"}
    if true {
        if true {
            var v4 uint16 = 42
            _ = v4
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v1 -= (v1 | v1)
    }
    v1 -= 32767
    var v5 uint16 = 65535
    _ = len(v2)
    switch int(v5) {
    case 128:
        v6 := GenS1{F0: "", F1: 4294967295}
        _ = v6
    case 32767:
        fmt.Println(v1, v1)
    }
    fmt.Println(v5, v1)
    v1 = v1
    if false {
        v7 := GenS1{F0: "foo bar", F1: 2}
        _ = v7
    }
    _ = len(v3)
    if len(v3) > 0 {
        v3[0] = ""
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        if (false || false) {
            v9 := GenS1{F0: ("" + "hello"), F1: 42}
            v10 := GenS1{F0: "foo bar", F1: 4294967295}
            _ = v10
            _ = v9
        }
        var v11 int8 = 1
        _ = fi8
        _ = v11
    }
    var v12 uint64 = 330
    var v13 uint64 = v12
    for fi14 := 0; fi14 < 8; fi14++ {
        if true {
            if len(v2) > 0 {
                v2[0] = fi14
            }
            v5 += (v5 % (((v5 + 255)) | 1))
            for fi15 := 0; fi15 < 5; fi15++ {
                v12 -= 42
                v5 -= v5
                var v16 uint = 1
                _ = fi15
                _ = v16
            }
        } else {
            for fi17 := 0; fi17 < 6; fi17++ {
                var v18 uint16 = (v5 << ((18446744073709551615) & 63))
                v12 = (v12 << ((728) & 63))
                v3 = append(v3, ("foo bar" + "\n\t"))
                var v19 uint64 = v13
                _ = fi17
                _ = v18
                _ = v19
            }
            var v20 uint32 = uint32(v12)
            _ = v20
        }
        _ = fi14
    }
    var v21 int32 = -351
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v3
    _ = v5
}
