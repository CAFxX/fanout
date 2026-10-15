package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 uint32 = 42
    var v2 bool = (((^(v1)) <= v1) || (1 > 66))
    var v3 int16 = 779
    v4 := make(map[int]uint32)
    var v5 bool = ((((960 >= 282) || false) || v2) && true)
    var v6 int8 = 1
    v2 = v2
    var v7 bool = (687 > 0)
    var v8 int8 = v6
    v9 := GenS1{F0: (v6 | (v6 & 12))}
    var v10 uint8 = 32
    var v11 int = -763
    if v5 {
        for fi12 := 0; fi12 < 5; fi12++ {
            fmt.Println(v10)
            _ = fi12
        }
        var v13 uint = 65535
        _ = v13
    } else {
        _ = len(v4)
    }
    v14 := []int{0}
    var v15 int64 = 42
    switch ("\n\t" + "a") {
    case "hello":
        switch (v11 & int(v8)) {
        case -1:
            var v16 int32 = 781
            _ = v16
        case 1:
            var v17 int64 = (v15 * 9223372036854775807)
            _ = v17
        case 529:
            v15 = (v15 - 1)
        default:
            for fi18 := 0; fi18 < 14; fi18++ {
                v10 += (^(v10))
                v14 = append(v14, (fi18 >> ((605) & 63)))
                if v5 {
                    var v19 bool = true
                    v8 += v8
                    _ = v19
                }
                _ = fi18
            }
        }
    case "\n\t":
        fmt.Println(v7)
    case "foo bar":
        v3 += (-(v3))
    }
    delete(v4, (v11 ^ (-(v11))))
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
