package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 uint8; F2 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{"\n\t"}
    var v3 int16 = -1
    v4 := make(map[string]int)
    var v5 int8 = 0
    v6 := []uint32{}
    var v7 int = 42
    if ((0.0 * -2.25) <= ((1.5 / 3.14159) / 0.1)) {
        delete(v4, "a")
        if (true && true) {
            if (-740 <= 42) {
                var v8 string = (("a" + "hello") + "\n\t")
                _ = v8
            }
            for fi9 := 0; fi9 < 2; fi9++ {
                for fi10 := 0; fi10 < 10; fi10++ {
                    fmt.Println(fi10, v7, fi9)
                    v11 := GenS1{F0: 40}
                    v12 := GenS2{F0: -341, F1: 1, F2: -393}
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                _ = fi9
            }
            v13 := GenS2{F0: 277, F1: 168, F2: 1}
            _ = v13
        }
    }
    if len(v6) > 0 {
        v6[0] = 1
    }
    var v14 uint64 = 4294967295
    var v15 uint32 = 1
    v16 := GenS3{F0: 723}
    v17 := GenS4{F0: 2147483647}
    for fi18 := 0; fi18 < 17; fi18++ {
        v7 += fi18
        _ = fi18
    }
    v2 = append(v2, (("hello" + "\n\t") + (("foo bar" + "hello") + "a")))
    v14 = (v14 * (v14 >> ((205) & 63)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
