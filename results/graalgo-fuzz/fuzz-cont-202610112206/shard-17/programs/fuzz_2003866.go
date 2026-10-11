package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int; F1 int }

func main() {
    var v1 float32 = ((0.1 - -2.25) - 1.5)
    var v2 string = ("" + (("" + ("" + "\n\t")) + "foo bar"))
    var v3 int32 = -2147483648
    var v4 uint32 = 2
    var v5 uint = 18446744073709551615
    var v6 uint32 = v4
    if false {
        fmt.Println(v6)
        switch ("a" + (v2 + v2)) {
        case "\n\t":
            var v7 uint = v5
            _ = v7
        default:
            v8 := GenS1{F0: 42}
            _ = v8
        }
    }
    v9 := GenS2{F0: -9223372036854775808}
    v10 := GenS3{F0: 4294967295}
    v2 = (v2 + ((v2 + v2) + ((v2 + "foo bar") + "foo bar")))
    v4 += v6
    fmt.Println(v3, v6)
    v11 := GenS4{F0: 616, F1: -166}
    fmt.Println(v1, v2)
    fmt.Println(v5, v4, v1)
    for fi12 := 0; fi12 < 10; fi12++ {
        v1 -= (float32(1.5) * (v1 / 3.14159))
        var v13 int32 = v3
        _ = fi12
        _ = v13
    }
    if (!(0.0 > v1)) {
        v14 := []int{32768, 13, 9223372036854775807, int(v3)}
        switch ((("a" + "a") + v2) + ("" + "foo bar")) {
        case "foo bar":
            v6 = v6
        case "":
            if ((v4 < v6) && (!(true || (1 < -316)))) {
                for fi15 := 0; fi15 < 9; fi15++ {
                    if (-32768 < 575) {
                        switch (fi15 - int(v5)) {
                        case 1:
                            fmt.Println(v5, v4)
                        }
                        var v16 int = fi15
                        var v17 uint16 = 0
                        _ = v16
                        _ = v17
                    }
                    _ = fi15
                }
                for fi18 := 0; fi18 < 4; fi18++ {
                    fmt.Println(v4)
                    _ = fi18
                }
            } else {
                var v19 uint64 = 598
                _ = v19
            }
        case "a":
            var v20 string = v2
            _ = v20
        default:
            fmt.Println(v1, v1)
        }
        var v21 int32 = (v3 * v3)
        _ = v14
        _ = v21
    }
    v3 = (v3 & v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
