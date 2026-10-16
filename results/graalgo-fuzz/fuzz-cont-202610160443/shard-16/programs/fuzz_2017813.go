package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 uint8; F1 int32; F2 uint64 }
type GenS3 struct { F0 int32; F1 int64 }

func main() {
    v1 := []string{"foo bar", "foo bar", ("" + (((("" + "\n\t") + ("\n\t" + "")) + "") + ("foo bar" + "a"))), ("" + "")}
    var v2 int16 = 0
    v3 := make(map[int]uint32)
    for fi4 := 0; fi4 < 17; fi4++ {
        v5 := GenS1{F0: 32767, F1: 536}
        var v6 float64 = 3.14159
        _ = fi4
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 4; fi7++ {
        for fi8 := 0; fi8 < 7; fi8++ {
            switch (("a" + ("a" + "foo bar")) + "") {
            case "a":
                var v9 int8 = -11
                _ = v9
            case "\n\t":
                _ = len(v3)
            case "foo bar":
                v2 += (v2 | v2)
            }
            v2 += (v2 >> ((482) & 63))
            var v10 bool = false
            _ = fi8
            _ = v10
        }
        _ = fi7
    }
    v2 += v2
    v2 += -733
    if false {
        for fi11 := 0; fi11 < 18; fi11++ {
            v12 := GenS1{F0: 255, F1: 578}
            if true {
                delete(v3, (fi11 * fi11))
                for fi13 := 0; fi13 < 19; fi13++ {
                    v2 -= (^(v2))
                    v14 := GenS1{F0: 0, F1: 2}
                    fmt.Println(v2, fi13)
                    _ = fi13
                    _ = v14
                }
                fmt.Println(fi11, v2)
            }
            _ = fi11
            _ = v12
        }
    } else {
        if (42 >= 4294967295) {
            v2 = v2
            for fi15 := 0; fi15 < 20; fi15++ {
                if true {
                    var v16 string = "a"
                    v2 = v2
                    v17 := GenS1{F0: 1, F1: 42}
                    _ = v16
                    _ = v17
                }
                v2 = 1
                v2 += v2
                _ = fi15
            }
            v18 := GenS1{F0: 42, F1: 2147483647}
            _ = v18
        } else {
            if (128 > 0) {
                v19 := GenS2{F0: 128, F1: 358, F2: 18446744073709551615}
                _ = v19
            }
            var v20 string = "hello"
            var v21 int64 = 989
            _ = v20
            _ = v21
        }
        var v22 int32 = 2147483647
        v23 := GenS2{F0: 127, F1: v22, F2: 4294967295}
        _ = v22
        _ = v23
    }
    delete(v3, 42)
    for fi24 := 0; fi24 < 5; fi24++ {
        v2 += v2
        _ = fi24
    }
    for fi25 := 0; fi25 < 5; fi25++ {
        v2 = v2
        v1 = append(v1, "a")
        if false {
            if false {
                var v26 uint16 = 2
                v2 -= (^(v2))
                v27 := GenS2{F0: 109, F1: 907, F2: 955}
                _ = v26
                _ = v27
            }
        }
        _ = fi25
    }
    var v28 uint8 = 119
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    _ = len(v1)
    v29 := GenS3{F0: 42, F1: 9223372036854775807}
    var v30 uint16 = 42
    var v31 int64 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
