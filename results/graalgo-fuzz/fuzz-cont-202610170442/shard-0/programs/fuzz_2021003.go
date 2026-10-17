package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 uint16 = 2
    var v2 uint = 251
    var v3 uint8 = uint8(v2)
    var v4 string = "foo bar"
    var v5 bool = (2 != 2147483647)
    var v6 uint16 = v1
    if v5 {
        for fi7 := 0; fi7 < 11; fi7++ {
            v2 = v2
            v8 := []int{}
            _ = fi7
            _ = v8
        }
        for fi9 := 0; fi9 < 4; fi9++ {
            var v10 int32 = 2147483647
            _ = fi9
            _ = v10
        }
    } else {
        switch -516 {
        case 127:
            switch 9223372036854775807 {
            case -1:
                for fi11 := 0; fi11 < 6; fi11++ {
                    for fi12 := 0; fi12 < 7; fi12++ {
                        v13 := []int{}
                        var v14 uint32 = 2
                        v13 = append(v13, (fi12 / ((fi12) | 1)))
                        v15 := GenS1{F0: 32768}
                        v16 := GenS1{F0: fi11}
                        _ = fi12
                        _ = v13
                        _ = v14
                        _ = v15
                        _ = v16
                    }
                    switch (v4 + v4) {
                    case "foo bar":
                        v17 := GenS1{F0: -2147483648}
                        _ = v17
                    }
                    _ = fi11
                }
            case 2147483647:
                v18 := GenS1{F0: -792}
                _ = v18
            }
        case -9223372036854775808:
            v1 = (v1 & uint16(v2))
        case -2147483648:
            if ((false || v5) || v5) {
                for fi19 := 0; fi19 < 8; fi19++ {
                    switch fi19 {
                    case 9223372036854775807:
                        var v20 bool = (693 != 2147483647)
                        _ = v20
                    case 2147483647:
                        v21 := GenS1{F0: fi19}
                        _ = v21
                    case 466:
                        v4 = v4
                    }
                    _ = fi19
                }
            }
        default:
            v4 = v4
        }
        v22 := GenS1{F0: -2147483648}
        _ = v22
    }
    v5 = false
    for fi23 := 0; fi23 < 16; fi23++ {
        v24 := GenS1{F0: fi23}
        _ = fi23
        _ = v24
    }
    v25 := GenS1{F0: 42}
    switch 127 {
    case 32768:
        var v26 uint = (v2 << ((v2) & 63))
        _ = v26
    case 32767:
        v27 := GenS1{F0: 32767}
        _ = v27
    default:
        v5 = (32767 > 1)
    }
    if ((v5 && (1 <= 4294967295)) || (v1 > (v6 << ((1) & 63)))) {
        switch v4 {
        case "foo bar":
            for fi28 := 0; fi28 < 7; fi28++ {
                v29 := []int{fi28}
                var v30 uint8 = v3
                _ = fi28
                _ = v29
                _ = v30
            }
        default:
            var v31 uint64 = 966
            _ = v31
        }
        v1 = (v6 - 1)
    }
    var v32 int32 = -2147483648
    var v33 int = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
