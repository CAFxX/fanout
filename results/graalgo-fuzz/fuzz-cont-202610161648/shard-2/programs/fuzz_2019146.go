package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 string, p1 uint16, p2 int8) uint {
    p0 = p0
    if true {
        for fi5 := 0; fi5 < 12; fi5++ {
            v6 := GenS1{F0: -563}
            p1 = (p1 >> ((18446744073709551615) & 63))
            for fi7 := 0; fi7 < 3; fi7++ {
                if true {
                    switch (fi5 - (fi5 - fi7)) {
                    case 490:
                        var v8 string = (("" + p0) + (p0 + p0))
                        _ = v8
                    case 0:
                        var v9 string = ((p0 + "foo bar") + (((p0 + p0) + "a") + p0))
                        _ = v9
                    }
                    p1 = 65535
                }
                var v10 int8 = p2
                _ = fi7
                _ = v10
            }
            _ = fi5
            _ = v6
        }
        for fi11 := 0; fi11 < 4; fi11++ {
            switch "a" {
            case "a":
                p2 = (p2 & 42)
            }
            _ = fi11
        }
        v12 := []int{0, 1}
        _ = v12
    }
    switch -9223372036854775808 {
    case 0:
        var v13 uint64 = 783
        _ = v13
    case -2147483648:
        var v14 uint16 = 795
        _ = v14
    case 42:
        p0 = ("foo bar" + "\n\t")
    }
    var v15 int64 = -1
    v16 := []int{650, 528}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    return 255
}

func main() {
    var v1 uint64 = 42
    var v2 string = ("\n\t" + ((("foo bar" + "a") + (("foo bar" + "a") + "\n\t")) + "hello"))
    var v3 int = 845
    var v4 uint32 = 0
    v17 := []int{675, v3}
    fmt.Println(v1)
    fmt.Println(v3, v3, v1)
    v1 = (v1 - (v1 / (((v1 & v1)) | 1)))
    fmt.Println(v1, v3, v4)
    v3 += v3
    var v18 uint16 = 126
    v4 = v4
    v19 := GenS1{F0: -165}
    v20 := GenS1{F0: 971}
    var v21 uint8 = 1
    v21 = (v21 | v21)
    fmt.Println(v2, v3, v2)
    fmt.Println(v18)
    if (644 <= 4294967295) {
        v22 := GenS1{F0: -637}
        for fi23 := 0; fi23 < 2; fi23++ {
            v24 := GenS1{F0: 758}
            v25 := GenS1{F0: -1}
            v3 = 42
            _ = fi23
            _ = v24
            _ = v25
        }
        _ = v22
    }
    v18 = 621
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
