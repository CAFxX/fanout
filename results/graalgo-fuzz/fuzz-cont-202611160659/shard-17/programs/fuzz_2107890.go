package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int16 }
type GenS2 struct { F0 string; F1 uint64; F2 uint32 }
func genFn0() uint16 {
    v6 := GenS1{F0: 969, F1: 15, F2: 570}
    v7 := GenS1{F0: 2, F1: 32768, F2: -32768}
    var v8 uint = 1
    _ = v6
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := []uint32{2147483647}
    var v2 string = ("a" + "hello")
    var v3 float64 = (1.5 / 1.0)
    v4 := []string{}
    v5 := make(map[uint32]bool)
    for fi9 := 0; fi9 < 19; fi9++ {
        for fi10 := 0; fi10 < 20; fi10++ {
            v3 = v3
            if (148 <= 255) {
                v5[42] = false
                switch ("\n\t" + "hello") {
                case "\n\t":
                    fmt.Println(v2)
                case "foo bar":
                    var v11 uint64 = 0
                    _ = v11
                case "a":
                    fmt.Println(fi10)
                }
            }
            _ = fi10
        }
        v12 := GenS1{F0: 324, F1: (-(fi9)), F2: -32768}
        _ = fi9
        _ = v12
    }
    var v13 uint8 = 232
    if ((("\n\t" + "hello") + v2) < (v2 + (v2 + (("foo bar" + ((v2 + v2) + (v2 + v2))) + "foo bar")))) {
        v2 = (v2 + (v2 + "a"))
    }
    var v14 uint = 2
    v15 := GenS1{F0: 255, F1: -125, F2: 1000}
    v14 += 65535
    v3 = v3
    for fi16 := 0; fi16 < 16; fi16++ {
        var v17 float32 = (1e10 * float32(v3))
        if (("\n\t" + v2) >= (v2 + "\n\t")) {
            if true {
                v3 -= (v3 * -2.25)
            }
            v18 := GenS2{F0: "foo bar", F1: 168, F2: 0}
            switch ((("a" + v2) + v2) + v2) {
            case "a":
                v2 = v2
            }
            _ = v18
        }
        for fi19 := 0; fi19 < 19; fi19++ {
            v13 = (v13 / (((v13 & v13)) | 1))
            delete(v5, 0)
            if (32767 >= fi16) {
                v17 = v17
            } else {
                var v20 float32 = -2.25
                v21 := GenS2{F0: v2, F1: 379, F2: 529}
                _ = len(v4)
                v1 = append(v1, 851)
                _ = v20
                _ = v21
            }
            _ = fi19
        }
        _ = fi16
        _ = v17
    }
    _ = len(v4)
    delete(v5, 1)
    var v22 float32 = 0.1
    fmt.Println(v3, v2, v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
