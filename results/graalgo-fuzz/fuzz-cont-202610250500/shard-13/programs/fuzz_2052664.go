package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }

func main() {
    var v1 uint16 = 2
    v2 := make(map[uint32]string)
    var v3 int = 4294967295
    v3 += 9223372036854775807
    v4 := []int{4294967295}
    if false {
        if (true && false) {
            if (true && (0 != 657)) {
                if (-147 >= 1) {
                    fmt.Println(v3, v1)
                    var v5 uint8 = 42
                    fmt.Println(v5, v3, v3)
                    _ = v5
                } else {
                    var v6 float64 = (float64(0.1) + 0.1)
                    _ = v6
                }
            }
            v1 -= (v1 % ((v1) | 1))
        }
        v7 := GenS1{F0: v1, F1: ("a" + "foo bar")}
        v3 -= -2147483648
        _ = v7
    }
    v8 := GenS1{F0: (^(v1)), F1: ("hello" + (("foo bar" + "") + "hello"))}
    v3 -= (v3 << ((730) & 63))
    var v9 float32 = 0.0
    var v10 float32 = float32(1.5)
    var v11 float64 = 1.5
    v12 := GenS1{F0: v1, F1: ""}
    _ = len(v2)
    if (!(false || ((-32768 >= -311) || true))) {
        v3 = v3
    } else {
        switch ("\n\t" + "hello") {
        case "hello":
            var v13 uint = 0
            _ = v13
        case "foo bar":
            var v14 bool = (80 == 255)
            _ = v14
        default:
            fmt.Println(v1)
        }
    }
    var v15 uint16 = v1
    var v16 uint16 = (v15 << ((365) & 63))
    v17 := GenS1{F0: 277, F1: ("\n\t" + (("a" + ("" + "foo bar")) + ("hello" + "hello")))}
    fmt.Println(v9, v16, v11)
    for fi18 := 0; fi18 < 17; fi18++ {
        v9 = (v9 - v9)
        switch "a" {
        case "":
            fmt.Println(v11, v9, v16)
        default:
            for fi19 := 0; fi19 < 1; fi19++ {
                switch (("a" + "\n\t") + ("" + "\n\t")) {
                case "hello":
                    v1 = (^(v15))
                default:
                    v20 := GenS1{F0: 613, F1: "a"}
                    _ = v20
                }
                _ = fi19
            }
        }
        v21 := GenS1{F0: v16, F1: "hello"}
        _ = fi18
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
