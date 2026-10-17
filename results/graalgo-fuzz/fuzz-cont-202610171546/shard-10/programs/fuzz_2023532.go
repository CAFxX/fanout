package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int }
func genFn0(p0 uint8) uint {
    for fi5 := 0; fi5 < 20; fi5++ {
        p0 += p0
        _ = fi5
    }
    v6 := []int{}
    for fi7 := 0; fi7 < 19; fi7++ {
        p0 = (p0 ^ (p0 * 147))
        _ = fi7
    }
    p0 = p0
    _ = p0
    _ = v6
    return 4294967295
}
func genFn1(p0 uint64) int8 {
    var v8 int8 = -86
    var v9 string = "a"
    _ = p0
    _ = v8
    _ = v9
    return v8
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = 0
    var v3 int8 = 127
    var v4 string = (("\n\t" + "foo bar") + ("\n\t" + "\n\t"))
    var v10 int8 = (v3 + genFn1(1))
    v1[2] = v4
    v4 = (v4 + v4)
    v2 += v2
    fmt.Println(v3, v4)
    var v11 uint64 = 18446744073709551615
    switch 32768 {
    case 42:
        if (2 < 42) {
            var v12 int = 0
            _ = v12
        } else {
            var v13 float32 = float32(-2.25)
            var v14 uint64 = (v11 & 4294967295)
            _ = v13
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 3; fi15++ {
        if true {
            for fi16 := 0; fi16 < 3; fi16++ {
                if (genFn1((^(v11))) == (v3 % ((genFn1(v11)) | 1))) {
                    v17 := GenS1{F0: 32767, F1: v11, F2: fi16}
                    v18 := GenS1{F0: fi15, F1: v11, F2: (fi15 ^ 960)}
                    v19 := []int{fi16, (^(fi16)), fi16}
                    v20 := GenS1{F0: fi15, F1: (^(v11)), F2: 569}
                    _ = v17
                    _ = v18
                    _ = v19
                    _ = v20
                }
                _ = fi16
            }
            for fi21 := 0; fi21 < 16; fi21++ {
                v10 += v10
                var v22 string = v4
                _ = fi21
                _ = v22
            }
        }
        v4 = v4
        _ = fi15
    }
    v4 = ((v4 + "hello") + ("foo bar" + v4))
    var v23 int16 = 359
    for fi24 := 0; fi24 < 5; fi24++ {
        v10 = (v10 - -1)
        if (1 <= 197) {
            for fi25 := 0; fi25 < 8; fi25++ {
                v26 := GenS1{F0: (fi24 ^ (fi24 ^ (fi25 / ((fi24) | 1)))), F1: 4294967295, F2: fi24}
                var v27 float32 = 1e10
                var v28 uint64 = v11
                _ = fi25
                _ = v26
                _ = v27
                _ = v28
            }
            var v29 uint = genFn0(0)
            var v30 float32 = (1.5 + 0.0)
            _ = v29
            _ = v30
        } else {
            switch v4 {
            case "\n\t":
                switch 1 {
                case 32768:
                    v11 = (v11 * 4294967295)
                case -9223372036854775808:
                    v31 := []int{}
                    _ = v31
                case 128:
                    fmt.Println(v2, v10, v11)
                }
            case "a":
                v32 := GenS1{F0: 2147483647, F1: v11, F2: -1}
                _ = v32
            }
            var v33 uint64 = (v11 >> ((genFn0(127)) & 63))
            var v34 uint32 = 0
            _ = v33
            _ = v34
        }
        _ = fi24
    }
    fmt.Println(v23, v23)
    var v35 uint16 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v23
    _ = v3
    _ = v35
    _ = v4
}
