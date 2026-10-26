package main

import (
    "fmt"
)


func main() {
    var v1 int = 32768
    var v2 uint32 = 4294967295
    v3 := make(map[uint32]string)
    var v4 uint8 = 30
    v5 := []uint32{}
    var v6 uint32 = (v2 >> ((255) & 63))
    if (float32(float64(-2.25)) <= float32((0.1 - 1.5))) {
        _ = len(v5)
        var v7 int = v1
        _ = v7
    } else {
        v6 -= 0
        _ = len(v3)
        if true {
            for fi8 := 0; fi8 < 17; fi8++ {
                switch (v1 % ((v1) | 1)) {
                case -549:
                    var v9 string = (("foo bar" + "\n\t") + "foo bar")
                    _ = v9
                case 586:
                    fmt.Println(v4, v1, v6)
                }
                if false {
                    v10 := v3[v6]
                    v1 -= fi8
                    _ = v10
                }
                _ = fi8
            }
            var v11 bool = false
            _ = v11
        } else {
            var v12 uint64 = 66
            for fi13 := 0; fi13 < 17; fi13++ {
                var v14 uint64 = (v12 << ((1) & 63))
                _ = fi13
                _ = v14
            }
            _ = v12
        }
    }
    v6 = v2
    v1 -= v1
    v6 += 65535
    var v15 bool = ("\n\t" != "foo bar")
    for fi16 := 0; fi16 < 9; fi16++ {
        var v17 uint16 = 42
        v4 += (v4 | v4)
        fmt.Println(v15)
        _ = fi16
        _ = v17
    }
    var v18 string = ""
    switch (v18 + ("" + v18)) {
    case "foo bar":
        v18 = ""
    case "hello":
        switch (((v18 + v18) + v18) + ("foo bar" + v18)) {
        case "\n\t":
            var v19 int32 = 2147483647
            _ = v19
        case "hello":
            v1 -= v1
        }
    case "":
        v6 = 2147483648
    }
    v4 += uint8(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
