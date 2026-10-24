package main

import (
    "fmt"
)

func genFn0() uint32 {
    if (false || false) {
        var v6 string = "a"
        var v7 int64 = 0
        _ = v6
        _ = v7
    }
    if (-783 == -202) {
        var v8 int8 = -1
        _ = v8
    }
    var v9 uint32 = 65535
    var v10 int64 = 42
    _ = v10
    _ = v9
    return (v9 * (v9 % (((v9 % (((v9 & (v9 >> ((2) & 63)))) | 1))) | 1)))
}
func genFn1(p0 bool) int16 {
    v11 := []int{0, 127, 4294967295}
    fmt.Println(p0)
    _ = p0
    _ = v11
    return 0
}

func main() {
    var v1 float32 = (0.1 * (0.0 * 0.0))
    var v2 float64 = 0.1
    var v3 int16 = 32767
    var v4 int16 = -32768
    v5 := make(map[string]uint32)
    var v12 uint8 = 2
    var v13 bool = true
    var v14 uint8 = (v12 << ((268) & 63))
    var v15 int = 954
    var v16 float64 = v2
    var v17 int16 = v3
    if v13 {
        v18 := []int{}
        _ = v18
    } else {
        fmt.Println(v16)
        if v13 {
            v4 = (^(v3))
            switch "" {
            case "foo bar":
                fmt.Println(v13, v15)
            default:
                if v13 {
                    v4 = genFn1(((2 < 0) && (2147483648 <= 65535)))
                    for fi19 := 0; fi19 < 14; fi19++ {
                        switch (^(fi19)) {
                        case 0:
                            v3 -= genFn1(false)
                        case 127:
                            var v20 float64 = 3.14159
                            _ = v20
                        case 128:
                            var v21 float32 = (v1 * 0.1)
                            _ = v21
                        }
                        _ = len(v5)
                        v1 += (3.14159 * 1e10)
                        _ = fi19
                    }
                } else {
                    fmt.Println(v1)
                }
            }
            var v22 int = (v15 % ((v15) | 1))
            _ = v22
        } else {
            var v23 float32 = v1
            _ = v23
        }
    }
    delete(v5, ((("\n\t" + ("foo bar" + ("\n\t" + ("hello" + "")))) + "foo bar") + "foo bar"))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
