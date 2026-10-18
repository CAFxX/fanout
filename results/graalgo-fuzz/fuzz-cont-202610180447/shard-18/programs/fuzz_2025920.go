package main

import (
    "fmt"
)

func genFn0(p0 uint16) uint {
    var v5 int64 = 9223372036854775807
    p0 -= (^(p0))
    _ = p0
    _ = v5
    return 4294967295
}

func main() {
    var v1 string = ("\n\t" + "foo bar")
    v2 := []string{"foo bar", "foo bar", "hello", v1}
    v3 := make(map[int]string)
    var v4 uint16 = 14
    fmt.Println(v1, v4)
    v4 += (v4 / (((v4 - (^(v4)))) | 1))
    _ = len(v2)
    switch -565 {
    case -787:
        fmt.Println(v1, v1)
    case 127:
        for fi6 := 0; fi6 < 9; fi6++ {
            for fi7 := 0; fi7 < 3; fi7++ {
                fmt.Println(fi6)
                _ = fi7
            }
            for fi8 := 0; fi8 < 12; fi8++ {
                v1 = (v1 + (("hello" + v1) + "\n\t"))
                _ = fi8
            }
            switch fi6 {
            case 1:
                if (1 <= 4294967295) {
                    for fi9 := 0; fi9 < 7; fi9++ {
                        _ = len(v2)
                        _ = fi9
                    }
                } else {
                    var v10 int64 = 0
                    var v11 int16 = -32768
                    switch (fi6 / (((fi6 >> ((680) & 63))) | 1)) {
                    case -492:
                        _ = len(v2)
                    }
                    _ = v10
                    _ = v11
                }
            case 42:
                if false {
                    if false {
                        v3[fi6] = v1
                    }
                }
            }
            _ = fi6
        }
    case 32767:
        if (true && (false && false)) {
            v4 += 2
        } else {
            delete(v3, 127)
            if false {
                var v12 int = 9223372036854775807
                v4 -= v4
                var v13 uint32 = 2147483647
                _ = v12
                _ = v13
            }
        }
    }
    var v14 uint16 = v4
    v14 = v14
    v1 = v1
    fmt.Println(v14, v4, v14)
    v14 = v4
    for fi15 := 0; fi15 < 4; fi15++ {
        v4 -= (^(v14))
        v14 -= (v14 % (((v4 * v4)) | 1))
        _ = fi15
    }
    v16 := v3[-2147483648]
    var v17 uint8 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
