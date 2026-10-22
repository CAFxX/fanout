package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 int32; F1 int16; F2 uint64 }

func main() {
    var v1 bool = ((!false) && (417 >= 4294967295))
    var v2 uint = 0
    var v3 uint64 = 368
    var v4 bool = true
    if (uint64(v2) > v3) {
        fmt.Println(v4, v1, v1)
        v3 = (v3 - v3)
    }
    v5 := GenS1{F0: 127, F1: 0}
    fmt.Println(v1, v3, v4)
    for fi6 := 0; fi6 < 16; fi6++ {
        switch "foo bar" {
        case "\n\t":
            v7 := []int{(fi6 * (^(fi6)))}
            _ = v7
        case "":
            if (!(-2.25 >= 3.14159)) {
                for fi8 := 0; fi8 < 8; fi8++ {
                    v9 := []int{}
                    if len(v9) > 0 {
                        v9[0] = (-(fi6))
                    }
                    _ = fi8
                    _ = v9
                }
                var v10 int = 42
                _ = v10
            }
        }
        _ = fi6
    }
    for fi11 := 0; fi11 < 2; fi11++ {
        switch (fi11 << (((v2 * v2)) & 63)) {
        case 128:
            if (!((2 > 0) && true)) {
                var v12 int = fi11
                _ = v12
            } else {
                var v13 float32 = (3.14159 + (-2.25 - float32((-2.25 * 1e10))))
                _ = v13
            }
        }
        _ = fi11
    }
    v14 := []int{}
    v3 = v3
    v4 = v4
    v1 = (42 <= 32767)
    v15 := GenS2{F0: -456, F1: -223, F2: (v3 - v3)}
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    v4 = true
    switch "" {
    case "a":
        v16 := GenS1{F0: -7, F1: -1}
        _ = v16
    }
    switch ((("\n\t" + "\n\t") + "") + ("" + (("" + ("a" + (("hello" + "a") + "a"))) + (("\n\t" + "a") + "hello")))) {
    case "hello":
        for fi17 := 0; fi17 < 9; fi17++ {
            for fi18 := 0; fi18 < 15; fi18++ {
                fmt.Println(v3, v4, v2)
                v4 = (v1 || true)
                _ = fi18
            }
            switch "" {
            case "":
                for fi19 := 0; fi19 < 9; fi19++ {
                    if true {
                        var v20 float32 = 3.14159
                        if len(v14) >= 2 {
                            v14 = v14[1:len(v14)]
                        }
                        _ = v20
                    } else {
                        var v21 uint32 = 364
                        v3 = (v3 << ((0) & 63))
                        fmt.Println(v21, v21, v2)
                        _ = v21
                    }
                    _ = fi19
                }
            default:
                fmt.Println(fi17, v3)
            }
            fmt.Println(fi17, v2)
            _ = fi17
        }
    case "a":
        v22 := GenS2{F0: 403, F1: 1, F2: v3}
        _ = v22
    case "":
        v2 -= v2
    default:
        v1 = v4
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
