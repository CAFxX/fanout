package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 uint; F2 string }

func main() {
    var v1 uint32 = 2147483648
    v2 := []uint32{(^(v1)), v1}
    var v3 uint64 = 0
    v4 := []string{("foo bar" + "hello"), "\n\t"}
    var v5 int = 32767
    var v6 int32 = 154
    var v7 int = (v5 << ((792) & 63))
    var v8 uint8 = 42
    v1 = v1
    if (0 <= (v3 - 4294967295)) {
        v9 := GenS1{F0: 1}
        var v10 uint32 = v1
        if (v6 < 953) {
            var v11 float32 = 1e10
            _ = v11
        } else {
            var v12 uint16 = 0
            var v13 uint32 = (v1 & (v1 - (^(v10))))
            _ = v12
            _ = v13
        }
        _ = v10
        _ = v9
    }
    var v14 uint8 = (v8 - v8)
    if true {
        if true {
            var v15 int32 = v6
            v6 -= (v6 ^ (v15 + (v15 + 0)))
            _ = v15
        }
        if ((v14 % (((v14 & 42)) | 1)) == (^(v14))) {
            var v16 bool = (false || ((-2.25 - 1.5) != (1.5 / (((((-2.25 - 1.5) * (((0.0 * -2.25) * 0.1) + 1e10)) * ((0.0 - ((((1.5 - float64(((1e10 + float32((0.0 * 0.1))) - (-2.25 + (0.0 / 1.5))))) + 1e10) / (0.1 / 1.0)) + 0.1)) + (1.5 * 1.5))) * 3.14159) / float64(-2.25)))))
            _ = v16
        } else {
            if len(v4) > 0 {
                v4[0] = ("" + "foo bar")
            }
            var v17 int32 = int32(v8)
            _ = v17
        }
        v18 := GenS2{F0: ("foo bar" + ""), F1: 0, F2: ("\n\t" + ("foo bar" + (("" + "foo bar") + (("a" + "a") + (("" + "a") + ("a" + "foo bar"))))))}
        _ = v18
    }
    fmt.Println(v8, v6, v8)
    var v19 bool = ("a" > "\n\t")
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
