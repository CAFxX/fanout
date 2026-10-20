package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32 }

func main() {
    var v1 int8 = 127
    v2 := make(map[int]bool)
    var v3 int64 = 9223372036854775807
    var v4 int16 = 42
    for fi5 := 0; fi5 < 5; fi5++ {
        for fi6 := 0; fi6 < 15; fi6++ {
            v7 := GenS1{F0: 2}
            var v8 bool = true
            v3 = 983
            if (!v8) {
                var v9 float32 = 0.0
                _ = v9
            }
            _ = fi6
            _ = v7
            _ = v8
        }
        _ = fi5
    }
    v1 = v1
    v4 += 27
    for fi10 := 0; fi10 < 10; fi10++ {
        for fi11 := 0; fi11 < 6; fi11++ {
            var v12 int = (fi10 + (fi10 * fi11))
            switch "" {
            case "a":
                fmt.Println(v4, v1, v12)
            case "":
                v1 = v1
            case "hello":
                fmt.Println(v1, v1, v4)
            }
            switch (v12 | 40) {
            case 127:
                v3 += (v3 >> ((2) & 63))
            }
            var v13 string = ("" + "foo bar")
            _ = fi11
            _ = v12
            _ = v13
        }
        if (-2.25 < -2.25) {
            v14 := v2[(fi10 * 0)]
            v15 := []int{(fi10 * fi10), 4294967295}
            delete(v2, fi10)
            _ = v14
            _ = v15
        } else {
            var v16 uint16 = 255
            var v17 int8 = (v1 & (v1 ^ -42))
            _ = v16
            _ = v17
        }
        if (true || (0 > 0)) {
            var v18 int8 = v1
            v3 = v3
            switch -957 {
            case -2147483648:
                for fi19 := 0; fi19 < 11; fi19++ {
                    var v20 string = "a"
                    v21 := GenS1{F0: 643}
                    v22 := []int{fi19, (fi19 / (((fi19 * int(v3))) | 1)), (fi10 / (((fi19 & (fi19 | (^(fi19))))) | 1)), fi19}
                    var v23 uint = 18446744073709551615
                    _ = fi19
                    _ = v20
                    _ = v21
                    _ = v22
                    _ = v23
                }
            }
            _ = v18
        } else {
            v24 := GenS2{F0: -1}
            v25 := GenS2{F0: -482}
            _ = v24
            _ = v25
        }
        _ = fi10
    }
    var v26 int8 = -119
    fmt.Println(v3, v26)
    var v27 int = 747
    var v28 int = -1
    v4 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
