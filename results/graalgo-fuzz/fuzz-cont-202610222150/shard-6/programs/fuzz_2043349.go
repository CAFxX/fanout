package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 int32 = -550
    var v2 string = "a"
    var v3 int16 = 117
    var v4 float64 = 1.5
    var v5 int64 = 142
    switch v2 {
    case "a":
        v1 = v1
    }
    fmt.Println(v3, v2)
    var v6 uint64 = 1
    v7 := []int{127, -464, -9223372036854775808}
    v8 := GenS1{F0: int8(v5)}
    v9 := GenS1{F0: 42}
    var v10 int64 = (v5 + v5)
    for fi11 := 0; fi11 < 18; fi11++ {
        v2 = v2
        for fi12 := 0; fi12 < 14; fi12++ {
            switch (fi11 << ((2) & 63)) {
            case 448:
                if false {
                    var v13 bool = (((true || (v3 > (-(v3)))) || true) || (float64(0.1) < (((1.5 * v4) - ((1e10 - -2.25) + v4)) + ((v4 - 0.1) / 3.14159))))
                    var v14 int16 = (-(v3))
                    v15 := GenS1{F0: 42}
                    _ = v13
                    _ = v14
                    _ = v15
                }
            case -9223372036854775808:
                _ = len(v7)
            }
            if ((fi11 - (^(fi11))) < fi12) {
                var v16 uint = 42
                _ = v16
            }
            v7 = append(v7, -2147483648)
            _ = fi12
        }
        _ = fi11
    }
    switch 1 {
    case 799:
        v5 = -567
    }
    for fi17 := 0; fi17 < 20; fi17++ {
        for fi18 := 0; fi18 < 5; fi18++ {
            if (0.0 <= 1e10) {
                if (65535 <= 42) {
                    v5 = (v10 % ((v5) | 1))
                }
            }
            var v19 float64 = float64(0.1)
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    var v20 float64 = (((float64((1e10 / 3.14159)) - (0.1 / (0.0 - v4))) * (v4 * (((3.14159 + 1.5) - v4) - 3.14159))) * v4)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
