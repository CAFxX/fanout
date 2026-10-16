package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 int32 = -310
    var v2 int = 42
    var v3 uint8 = 127
    v3 += (v3 | v3)
    switch "foo bar" {
    case "hello":
        if ((657 != 2147483647) && true) {
            v2 += v2
            for fi4 := 0; fi4 < 12; fi4++ {
                for fi5 := 0; fi5 < 18; fi5++ {
                    var v6 float32 = -2.25
                    v7 := []int{fi4, (fi4 | fi5), (v2 % ((32767) | 1))}
                    _ = fi5
                    _ = v6
                    _ = v7
                }
                v2 = 127
                var v8 uint32 = 968
                _ = fi4
                _ = v8
            }
        }
    default:
        var v9 int8 = 100
        _ = v9
    }
    v2 -= 32767
    var v10 int32 = 987
    v10 -= 0
    fmt.Println(v10, v2, v2)
    v3 = (v3 << ((65535) & 63))
    var v11 uint = 65535
    for fi12 := 0; fi12 < 16; fi12++ {
        for fi13 := 0; fi13 < 12; fi13++ {
            switch (^(fi13)) {
            case 128:
                v3 = v3
            case 0:
                fmt.Println(v1)
            case 494:
                var v14 uint16 = 0
                _ = v14
            }
            v2 = 128
            v15 := GenS1{F0: 1}
            _ = fi13
            _ = v15
        }
        v16 := GenS1{F0: 65535}
        _ = fi12
        _ = v16
    }
    for fi17 := 0; fi17 < 6; fi17++ {
        for fi18 := 0; fi18 < 11; fi18++ {
            var v19 uint8 = v3
            v20 := []int{v2, v2}
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = fi17
    }
    if false {
        v11 -= v11
        var v21 uint64 = 46
        v1 += int32(v3)
        _ = v21
    } else {
        fmt.Println(v3, v1, v2)
    }
    var v22 int = 32767
    v23 := []int{}
    var v24 uint32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
