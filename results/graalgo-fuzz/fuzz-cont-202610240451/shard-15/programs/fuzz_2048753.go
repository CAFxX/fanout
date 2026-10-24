package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 uint32 }

func main() {
    var v1 int32 = 20
    v2 := []uint32{673, 4294967295, 93, 4294967295}
    var v3 uint32 = 65535
    v2 = append(v2, 721)
    v4 := GenS1{F0: "a", F1: "\n\t"}
    switch 4294967295 {
    case 9223372036854775807:
        for fi5 := 0; fi5 < 7; fi5++ {
            var v6 float32 = 1e10
            _ = fi5
            _ = v6
        }
    case 1:
        var v7 uint32 = (v3 - v3)
        _ = v7
    case -4:
        v3 += 979
    }
    v1 = v1
    switch ("a" + "hello") {
    case "":
        v8 := GenS2{F0: 0}
        _ = v8
    }
    for fi9 := 0; fi9 < 19; fi9++ {
        fmt.Println(fi9, fi9)
        v3 -= v3
        _ = fi9
    }
    v3 -= 483
    v10 := GenS3{F0: 127, F1: 4294967295}
    for fi11 := 0; fi11 < 1; fi11++ {
        v3 += (v3 + v3)
        _ = fi11
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        fmt.Println(v1)
        v3 += v3
        v13 := GenS1{F0: ("\n\t" + "foo bar"), F1: "a"}
        _ = fi12
        _ = v13
    }
    if ((!(!(true && (!(169 <= 1))))) && false) {
        v3 = v3
        v14 := GenS4{F0: v3}
        v3 -= 4294967295
        _ = v14
    }
    for fi15 := 0; fi15 < 9; fi15++ {
        fmt.Println(v3, v3, v3)
        v3 = v3
        _ = fi15
    }
    if false {
        var v16 int = 467
        v1 += v1
        _ = v16
    }
    fmt.Println(v3, v1)
    switch 4294967295 {
    case 1:
        if true {
            for fi17 := 0; fi17 < 14; fi17++ {
                v1 += v1
                v18 := GenS4{F0: 4294967295}
                if ((!((-186 >= int(v1)) && (136 == 2))) || false) {
                    fmt.Println(v3, fi17)
                }
                _ = fi17
                _ = v18
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if (0 >= 1) {
                var v19 uint64 = 4294967295
                var v20 bool = (true && true)
                _ = v19
                _ = v20
            } else {
                v3 = v3
                v2 = append(v2, v3)
            }
        }
    case -994:
        v2 = append(v2, v3)
    case -168:
        fmt.Println(v3)
    }
    var v21 int16 = 1000
    v22 := GenS3{F0: 255, F1: (v3 << ((2) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
