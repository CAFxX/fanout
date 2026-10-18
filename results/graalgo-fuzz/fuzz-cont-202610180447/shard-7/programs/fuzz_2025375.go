package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 uint8; F2 int8 }
type GenS4 struct { F0 uint; F1 int64 }

func main() {
    var v1 int64 = 0
    v2 := make(map[int]int)
    var v3 int32 = 2147483647
    v4 := []string{("" + ""), "hello"}
    v3 += 1
    if (677 <= 2) {
        for fi5 := 0; fi5 < 7; fi5++ {
            for fi6 := 0; fi6 < 10; fi6++ {
                if (false || true) {
                    delete(v2, (^(fi6)))
                    v4 = append(v4, ("" + ((("foo bar" + "") + "\n\t") + "hello")))
                    v3 = -570
                } else {
                    v7 := GenS1{F0: v1, F1: (!(1000 > 32767)), F2: -118}
                    var v8 int = (^(fi6))
                    var v9 uint16 = 32767
                    _ = v7
                    _ = v8
                    _ = v9
                }
                var v10 uint32 = 42
                fmt.Println(v3, v3, fi6)
                _ = fi6
                _ = v10
            }
            _ = fi5
        }
        v11 := v2[1]
        _ = v11
    }
    v2[128] = 0
    v12 := GenS1{F0: (^(v1)), F1: (!true), F2: -509}
    var v13 uint16 = 42
    switch 396 {
    case 689:
        switch (("" + "") + "") {
        case "a":
            var v14 uint = 255
            _ = v14
        case "hello":
            if (v3 >= v3) {
                for fi15 := 0; fi15 < 20; fi15++ {
                    v4 = append(v4, ("\n\t" + "a"))
                    v4 = append(v4, "")
                    _ = fi15
                }
            }
        case "foo bar":
            fmt.Println(v3, v1, v13)
        }
    case -2147483648:
        v16 := GenS2{F0: false}
        _ = v16
    }
    v17 := GenS3{F0: 319, F1: 255, F2: 71}
    if false {
        v13 = (v13 * (v13 + (^(v13))))
        var v18 int8 = -128
        var v19 int8 = (-(v18))
        _ = v18
        _ = v19
    } else {
        v13 = (^(v13))
        v4 = append(v4, "foo bar")
        v20 := GenS4{F0: 443, F1: -413}
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
