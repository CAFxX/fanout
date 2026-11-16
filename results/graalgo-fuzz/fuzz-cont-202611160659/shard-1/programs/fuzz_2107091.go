package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 uint64 }
type GenS2 struct { F0 uint8; F1 uint32; F2 int16 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]uint32)
    v3 := []string{}
    v4 := []int{42}
    var v5 int8 = -1
    v6 := GenS1{F0: 42, F1: 1, F2: 0}
    delete(v2, 2)
    v5 += (v5 & (^(v5)))
    v7 := GenS1{F0: 828, F1: 1, F2: 2}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v5 += v5
    switch 4294967295 {
    case 42:
        if (2 == 152) {
            switch (("foo bar" + ("foo bar" + (("foo bar" + ("" + "")) + "a"))) + ("" + ("hello" + "\n\t"))) {
            case "foo bar":
                v5 += v5
            case "a":
                v8 := GenS1{F0: 255, F1: 1, F2: 2}
                _ = v8
            }
            var v9 string = ("" + "a")
            _ = v9
        } else {
            if len(v3) > 0 {
                v3[0] = "a"
            }
            v5 -= (v5 - (v5 >> ((4294967295) & 63)))
            v5 = v5
        }
    case -134:
        v5 -= v5
    }
    if true {
        if false {
            var v10 float32 = 0.1
            delete(v1, ("\n\t" + "a"))
            var v11 uint8 = 42
            _ = v10
            _ = v11
        }
        for fi12 := 0; fi12 < 10; fi12++ {
            fmt.Println(fi12, fi12)
            v5 = v5
            _ = fi12
        }
        var v13 string = "a"
        _ = v13
    }
    var v14 int64 = 764
    var v15 int16 = -1
    var v16 uint32 = 923
    switch 0 {
    case -9223372036854775808:
        var v17 uint8 = 0
        _ = v17
    case 127:
        if ((false || true) || (true && (2 < 1))) {
            for fi18 := 0; fi18 < 11; fi18++ {
                fmt.Println(v5, fi18)
                for fi19 := 0; fi19 < 14; fi19++ {
                    delete(v2, (v16 >> ((957) & 63)))
                    v5 -= v5
                    v20 := GenS2{F0: 127, F1: 615, F2: (v15 / ((-495) | 1))}
                    _ = len(v1)
                    _ = fi19
                    _ = v20
                }
                fmt.Println(v15)
                _ = fi18
            }
            v2[v16] = v16
        } else {
            v15 = 32767
        }
    }
    if (!false) {
        if (!(0.1 < 0.1)) {
            fmt.Println(v16)
        }
    }
    v15 += -32768
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
