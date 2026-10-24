package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 bool }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 int32 }

func main() {
    var v1 int16 = -800
    var v2 float32 = ((3.14159 + 1e10) + 0.1)
    var v3 int = -881
    var v4 bool = true
    v5 := make(map[string]int)
    var v6 uint = 330
    v7 := GenS1{F0: 42, F1: 282, F2: (-447 != 9223372036854775807)}
    if false {
        var v8 uint = v6
        fmt.Println(v8)
        v1 = v1
        _ = v8
    } else {
        v9 := GenS1{F0: 4294967295, F1: 0, F2: v4}
        fmt.Println(v1, v3)
        _ = v9
    }
    v10 := GenS1{F0: 479, F1: 1, F2: (724 > 255)}
    fmt.Println(v3, v6)
    v11 := []int{(v3 << ((844) & 63)), (v3 * (v3 << ((v6) & 63)))}
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    switch ("\n\t" + "foo bar") {
    case "hello":
        if v4 {
            var v12 int8 = -1
            v2 -= 3.14159
            _ = v12
        }
    }
    v11 = append(v11, (v3 << ((v6) & 63)))
    v13 := GenS2{F0: 2147483648, F1: -2147483648}
    v14 := GenS3{F0: -682}
    for fi15 := 0; fi15 < 7; fi15++ {
        var v16 float32 = v2
        if (-81 >= 86) {
            if v4 {
                var v17 int32 = -2147483648
                _ = v17
            }
        } else {
            for fi18 := 0; fi18 < 19; fi18++ {
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                switch (fi15 >> ((18446744073709551615) & 63)) {
                case 2147483647:
                    switch (((("" + "") + "foo bar") + ("\n\t" + "hello")) + "") {
                    case "":
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                    case "a":
                        var v19 int32 = 1
                        _ = v19
                    }
                default:
                    v6 -= (v6 | (^(v6)))
                }
                _ = fi18
            }
            var v20 int16 = (v1 + v1)
            var v21 uint = (^(v6))
            _ = v20
            _ = v21
        }
        switch (fi15 / ((fi15) | 1)) {
        case 0:
            fmt.Println(v6)
        }
        _ = fi15
        _ = v16
    }
    delete(v5, "\n\t")
    var v22 int = -338
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
