package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64 }

func main() {
    var v1 int16 = 650
    v2 := make(map[int]int)
    v3 := make(map[string]int)
    var v4 int32 = 691
    v4 -= (^(v4))
    v4 += (v4 << ((18446744073709551615) & 63))
    for fi5 := 0; fi5 < 6; fi5++ {
        if true {
            v6 := v3["a"]
            v7 := []int{fi5, -1, (v6 % ((v6) | 1)), v6}
            switch (("hello" + "hello") + "foo bar") {
            case "a":
                var v8 int32 = (v4 + (v4 * v4))
                _ = v8
            case "\n\t":
                v9 := GenS1{F0: "foo bar", F1: 891, F2: 270}
                _ = v9
            case "hello":
                for fi10 := 0; fi10 < 8; fi10++ {
                    v6 -= (v6 << ((18446744073709551615) & 63))
                    if true {
                        v6 = fi5
                        fmt.Println(fi10)
                    }
                    _ = fi10
                }
            }
            _ = v6
            _ = v7
        } else {
            v4 = (v4 % ((-1) | 1))
        }
        _ = fi5
    }
    switch 127 {
    case 9223372036854775807:
        if ((-2.25 * (float64(float32((float64(0.0) + 1e10))) + 0.1)) >= 3.14159) {
            var v11 uint16 = 0
            var v12 int32 = v4
            _ = v11
            _ = v12
        }
    }
    fmt.Println(v4)
    var v13 uint64 = 42
    _ = len(v2)
    v13 += v13
    v13 = v13
    switch ((("foo bar" + ("hello" + ("\n\t" + ""))) + "\n\t") + "foo bar") {
    case "foo bar":
        for fi14 := 0; fi14 < 11; fi14++ {
            v15 := GenS2{F0: 255}
            v2[(fi14 & fi14)] = int(v1)
            v16 := GenS3{F0: 9223372036854775807}
            _ = fi14
            _ = v15
            _ = v16
        }
    case "":
        var v17 int16 = (v1 + 1)
        _ = v17
    }
    v18 := GenS2{F0: 255}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
