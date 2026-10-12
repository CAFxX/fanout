package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 int32 }
type GenS2 struct { F0 uint32; F1 int }

func main() {
    var v1 int8 = 0
    var v2 int64 = 42
    var v3 int32 = int32(v2)
    fmt.Println(v2, v2)
    v1 = (v1 + (-(v1)))
    var v4 uint32 = 0
    v5 := []int{32767}
    if (!true) {
        for fi6 := 0; fi6 < 20; fi6++ {
            v5 = append(v5, 32768)
            for fi7 := 0; fi7 < 15; fi7++ {
                var v8 string = ("foo bar" + ("foo bar" + "\n\t"))
                var v9 uint8 = 127
                _ = fi7
                _ = v8
                _ = v9
            }
            _ = fi6
        }
    } else {
        for fi10 := 0; fi10 < 5; fi10++ {
            v11 := GenS1{F0: (^(v4)), F1: 0, F2: (v3 + v3)}
            v12 := GenS1{F0: 2147483648, F1: 4294967295, F2: (v3 + 0)}
            _ = fi10
            _ = v11
            _ = v12
        }
        var v13 uint32 = (v4 >> ((65535) & 63))
        if ((v1 / ((v1) | 1)) == int8(v13)) {
            v14 := GenS1{F0: 2147483647, F1: 65535, F2: 42}
            for fi15 := 0; fi15 < 14; fi15++ {
                v16 := GenS1{F0: v13, F1: 17, F2: 1}
                _ = fi15
                _ = v16
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = v14
        } else {
            v2 += (v2 << ((42) & 63))
            v17 := GenS2{F0: 219, F1: -9223372036854775808}
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = v17
        }
        _ = v13
    }
    v1 = v1
    for fi18 := 0; fi18 < 11; fi18++ {
        v3 = v3
        fmt.Println(v1, v2, fi18)
        switch (fi18 + (-(fi18))) {
        case 529:
            v19 := GenS2{F0: v4, F1: fi18}
            _ = v19
        case 42:
            switch 861 {
            case -2147483648:
                v3 = (v3 ^ v3)
            default:
                var v20 bool = false
                _ = v20
            }
        }
        _ = fi18
    }
    v2 = (v2 ^ -1)
    if (false && true) {
        v4 = (v4 << ((65535) & 63))
        var v21 uint16 = 32767
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
