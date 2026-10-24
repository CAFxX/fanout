package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 float32 = -2.25
    var v2 uint16 = 255
    var v3 int64 = 42
    v4 := []int{42, 4294967295, 42}
    for fi5 := 0; fi5 < 15; fi5++ {
        fmt.Println(v1, fi5, v3)
        _ = fi5
    }
    var v6 string = "foo bar"
    v4 = append(v4, 343)
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 int8 = 42
        if len(v4) > 0 {
            v4[0] = 4294967295
        }
        _ = fi7
        _ = v8
    }
    v3 = (v3 << ((42) & 63))
    var v9 uint = 18446744073709551615
    v3 = (v3 >> ((v9) & 63))
    var v10 uint16 = v2
    for fi11 := 0; fi11 < 10; fi11++ {
        switch -611 {
        case -1:
            var v12 uint8 = 2
            _ = v12
        case 128:
            for fi13 := 0; fi13 < 20; fi13++ {
                if (!((1 <= (^(v3))) || (2147483648 > 42))) {
                    var v14 int16 = 0
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = v14
                } else {
                    v1 = v1
                }
                v15 := GenS1{F0: 128}
                _ = fi13
                _ = v15
            }
        case 807:
            v16 := GenS1{F0: 127}
            _ = v16
        }
        v9 = 128
        _ = fi11
    }
    for fi17 := 0; fi17 < 15; fi17++ {
        if (897 >= 42) {
            switch (fi17 << (((v9 >> ((18446744073709551615) & 63))) & 63)) {
            case -9223372036854775808:
                var v18 int8 = -29
                _ = v18
            case 1:
                fmt.Println(v9)
            }
            var v19 int8 = 64
            _ = v19
        }
        fmt.Println(v3)
        _ = len(v4)
        _ = fi17
    }
    v9 = (v9 >> (((v9 << (((v9 % ((v9) | 1))) & 63))) & 63))
    v2 -= v10
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v9
}
